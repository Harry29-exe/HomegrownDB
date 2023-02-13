package rawtype

import "HomegrownDB/lib/bparse"

type Args struct {
	Length   int
	Nullable bool
	VarLen   bool
	UTF8     bool
}

func SerializeArgs(args Args, s *bparse.Serializer) {
	s.Uint64(uint64(args.Length))
	s.Bool(args.Nullable)
	s.Bool(args.VarLen)
	s.Bool(args.UTF8)
}

func DeserializeArgs(d *bparse.Deserializer) Args {
	return Args{
		Length:   int(d.Uint64()),
		Nullable: d.Bool(),
		VarLen:   d.Bool(),
		UTF8:     d.Bool(),
	}
}
