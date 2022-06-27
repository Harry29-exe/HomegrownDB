import os
import subprocess

HOMEGROWN_DB_HOME = 'HOMEGROWN_DB_HOME'


def main():
    db_path = read_db_path()
    create_dir_hierarchy(db_path)
    export_db_env(db_path)

    print("success")


def read_db_path() -> str:
    print("Choose location for db storage [~/.homegrown_db]: ", end="")
    homegrown_db_path = input()
    if homegrown_db_path == '':
        home = os.getenv("HOME")
        homegrown_db_path = home + "/.homegrown_db"
    print("\t" + homegrown_db_path)

    return homegrown_db_path


def export_db_env(db_path: str):
    outfile = open(os.path.expanduser("~/.profile"), "a")

    if not "export " + HOMEGROWN_DB_HOME + "=" in outfile.read():
        outfile.write(
            "# homegrown db environment\n"
            "export " + HOMEGROWN_DB_HOME + "=" + db_path + "\n")
        subprocess.run("export " + HOMEGROWN_DB_HOME+"="+db_path)

    else:
        print("Environment variable: " + HOMEGROWN_DB_HOME + "is already set")
        quit(1)


def create_dir_hierarchy(db_path: str):
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
