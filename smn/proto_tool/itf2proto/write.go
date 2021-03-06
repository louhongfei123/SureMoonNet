package itf2proto

import (
	"fmt"
	"log"
	"strings"

	"github.com/ProtossGenius/SureMoonNet/basis/smn_file"
	"github.com/ProtossGenius/SureMoonNet/basis/smn_muti_write_cache"
	"github.com/ProtossGenius/SureMoonNet/basis/smn_pglang"
	"github.com/ProtossGenius/SureMoonNet/basis/smn_str"
)

func WriteProto(outDir string, list []*smn_pglang.ItfDef) error {
	pkg := list[0].Package
	oPath := outDir + "/rip_" + pkg + ".proto"
	file, err := smn_file.CreateNewFile(oPath)
	if err != nil {
		return err
	}
	defer file.Close()
	w := smn_muti_write_cache.NewFileMutiWriteCache()
	writeLine := func(f string, a ...interface{}) {
		w.WriteTail(fmt.Sprintf(f, a...) + "\n")
	}
	w.Append(smn_muti_write_cache.NewStrCache(fmt.Sprintf("syntax = \"proto3\";\noption java_package = \"pb\";\noption java_outer_classname=\"%s\";\npackage rip_%s;\n", pkg, pkg)))
	impts := smn_muti_write_cache.NewStrCache()
	impMap := make(map[string]bool)
	checkImport := func(typ string) {
		if typ == "net.Conn" {
			return
		}
		nameList := strings.Split(typ, ".")
		if len(nameList) == 1 {
			return
		}
		if !impMap[nameList[0]] {
			impts.WriteTail(fmt.Sprintf("import \"%s.proto\";", nameList[0]) + "\n")
			impMap[nameList[0]] = true
		}
	}
	w.Append(impts)
	for _, itf := range list {
		for _, mtd := range itf.Functions {
			if mtd.Name[0] < 'A' || mtd.Name[0] > 'Z' {
				log.Printf("warning! mtd.Name %s's first letter not upper\n", mtd.Name)
				continue
			}
			writeLine("message %s_%s_Prm {", itf.Name, mtd.Name)
			for i, prm := range mtd.Params {
				isArray, typ := smn_str.ProtoUseDeal(prm.Type)
				if typ == "net.Conn" {
					continue
				}
				checkImport(typ)
				rpt := ""
				if isArray {
					rpt = "repeated "
				}
				writeLine("\t%s%s %s = %d;", rpt, typ, prm.Var, i+1)
			}
			writeLine("}")
			writeLine("message %s_%s_Ret {", itf.Name, mtd.Name)
			for i, res := range mtd.Returns {
				isArray, typ := smn_str.ProtoUseDeal(res.Type)
				checkImport(typ)
				rpt := ""
				if isArray {
					rpt = "repeated "
				}
				writeLine("\t%s%s %s = %d;", rpt, typ, res.Var, i+1)
			}
			writeLine("}")
		}
	}
	w.Output(file)
	return nil
}
