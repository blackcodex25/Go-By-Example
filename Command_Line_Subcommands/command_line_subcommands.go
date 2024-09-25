package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}

/*  โปรแกรมนี้แสดงวิธีการใช้งาน subcommands บน command-line โดยใช้แพ็กเกจ flag ใน Go
1.การสร้างคำสั่งย่อย (Subcommands):
มีสอง subcommands: foo และ bar สร้างด้วย flag.NewFlagSet
fooCmd มี flags คือ enable (เป็น boolean) และ name (string)
barCmd มี flag คือ level (integer)
2.การตรวจสอบและแยกคำสั่ง:
ตรวจสอบ os.Args[1] เพื่อตรวจว่าผู้ใช้ระบุคำสั่งย่อยเป็น foo หรือ bar
3.การประมวลผล:
เมื่อเรียกคำสั่ง foo หรือ bar โปรแกรมจะแสดงผลค่าของ flag ตามที่ระบุ
*/
/* Subcommands คือ คำสั่งย่อยที่ใช้ในโปรแกรม command-line */
/* เพื่อจัดการคำสั่งที่ซับซ้อนหรือมีหลายฟังก์ชันย่อย โดยโปรแกรม */
/* จะแบ่งคำสั่งหลักออกเป็นชุดย่อยๆ และจัดการแต่ละคำสั่งย่อยแยกกัน */
/* เช่น คำสั่ง git มี subcommands หลายตัว เช่น git clone */
/* หรือ git commit ใน Go เราสามารถใช้แพ็กเกจ flag เพื่อสร้าง subcommands เหล่านี้ */
/* โดยแต่ละ subcommand จะมีชุด flag ที่เป็นของตนเอง */
