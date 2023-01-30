package rawtype

import "HomegrownDB/common/bparse"

type Args struct {
	Length   uint32
	Nullable bool
	VarLen   bool
	UTF8     bool
}

func SerializeArgs(args Args, s *bparse.Serializer) {
	s.Uint32(args.Length)
	s.Bool(args.Nullable)
	s.Bool(args.VarLen)
	s.Bool(args.UTF8)
}

func DeserializeArgs(d *bparse.Deserializer) Args {
	return Args{
		Length:   d.Uint32(),
		Nullable: d.Bool(),
		VarLen:   d.Bool(),
		UTF8:     d.Bool(),
	}
}
