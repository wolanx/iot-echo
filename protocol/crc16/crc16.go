package crc16

type Params struct {
	Poly   uint16
	Init   uint16
	RefIn  bool
	RefOut bool
	XorOut uint16
	Check  uint16
	Name   string
}

//CRC16_MODBUS  = Params{0x8005, 0xFFFF, true, true, 0x0000, 0x4B37, "CRC-16/MODBUS"}
//CRC16_GENIBUS = Params{0x1021, 0xFFFF, false, false, 0xFFFF, 0xD64E, "CRC-16/GENIUS"}

type Table struct {
	params Params
	data   [256]uint16
}

func MakeTable() *Table {
	table := new(Table)
	table.params = Params{0x1021, 0xFFFF, false, false, 0xFFFF, 0xD64E, "CRC-16/GENIBUS"}
	n := 0
	for j := 0; j <= 31; j++ {
		for i := 8 * j; i < 8*j+8; i++ {
			table.data[n] = crcEach(uint16(i))
			n++
		}
	}
	return table
}

func crcEach(data uint16) uint16 {
	var accum uint16 = 0x00

	data <<= 8
	for i := 8; i > 0; i-- {
		tmp := (data ^ accum) & 0x8000
		if tmp != 0 {
			accum = (accum << 1) ^ 0x1021
		} else {
			accum <<= 1
		}
		data <<= 1
	}

	return accum
}

func Init(table *Table) uint16 {
	return table.params.Init
}

func Update(accum uint16, arr []byte, table *Table) uint16 {
	accum = 0xffff
	for _, data := range arr {
		accum = (accum << 8) ^ table.data[byte(accum>>8)^data]
		//comb_val := (accum >> 8) ^ uint16(data)
		//tmp := crcEach(comb_val)
		//accum = tmp ^ (accum << 8)
	}
	return accum
}

func Checksum(data []byte, table *Table) uint16 {
	crc := Init(table)
	crc = Update(crc, data, table)
	return crc
}
