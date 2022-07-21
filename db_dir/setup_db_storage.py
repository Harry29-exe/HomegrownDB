import os
import shutil
import subprocess

HOMEGROWN_DB_HOME = 'HOMEGROWN_DB_HOME'
override_current_dir = False


def main():
    read_config()
    db_path = read_db_path()
    create_dir_hierarchy(db_path)
    export_db_env(db_path)

    print("success")


def read_config():
    print("Override current setup if exist? [n]")
    override = '_'
    while not (override != 't' or override != 'n' or override == ''):
        override = input()
        global override_current_dir

        if override == '' or override == 'n':
            override_current_dir = False
        elif override == 't':
            override_current_dir = True
        else:
            print('Please enter y or n')


def read_db_path() -> str:
    print("Choose location for db storage [~/.homegrown_db]: ", end="")
    homegrown_db_path = input()
    if homegrown_db_path == '':
        home = os.getenv("HOME")
        homegrown_db_path = home + "/.homegrown_db"
    print("\t" + homegrown_db_path)

    return homegrown_db_path


Z_PROFILE_PATH = "~/.zprofile"
BASH_PROFILE = "~/.bash_profile"


def export_db_env(db_path: str):
    path = os.path.expanduser(Z_PROFILE_PATH)
    if not os.path.exists(path):
        path = os.path.expanduser(BASH_PROFILE)

    outfile = open(path, "r+t")

    env_start = "export " + HOMEGROWN_DB_HOME + "="
    if not (env_start in outfile.read()):
        outfile.write(
            "# homegrown db environment\n"
            "export " + HOMEGROWN_DB_HOME + "=\"" + db_path + "\"\n")
        subprocess.run("export " + HOMEGROWN_DB_HOME + "=\"" + db_path + "\"", shell=True)

    else:
        print("Environment variable: " + HOMEGROWN_DB_HOME + " is already set")
        quit(1)


def create_dir_hierarchy(db_path: str):
    try:
        os.mkdir(db_path)
        dirs_to_create = [
            '/tables',
            '/lob',
            '/info',
            '/config',
        ]
        for dir_name in dirs_to_create:
            os.mkdir(db_path + dir_name)
    except FileExistsError:

        if not override_current_dir:
            print("\n\nOverride existing configuration is not enabled")
            quit(1)

        shutil.rmtree(db_path)
        os.mkdir(db_path)
        dirs_to_create = [
            '/tables',
            '/lob',
            '/info',
            '/config',
        ]
        for dir_name in dirs_to_create:
            os.mkdir(db_path + dir_name)


if __name__ == "__main__":
    main()
