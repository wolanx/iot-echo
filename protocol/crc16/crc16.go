package crc16

// Params represents parameters of CRC-16 algorithms.
// More information about algorithms parametrization and parameter descriptions
// can be found here - http://www.zlib.net/crc_v3.txt
type Params struct {
	Poly   uint16
	Init   uint16
	RefIn  bool
	RefOut bool
	XorOut uint16
	Check  uint16
	Name   string
}

// Predefined CRC-16 algorithms.
// List of algorithms with their parameters borrowed from here -  http://reveng.sourceforge.net/crc-catalogue/16.htm
//
// The variables can be used to create Table for the selected algorithm.
var (
	CRC16_ARC         = Params{0x8005, 0x0000, true, true, 0x0000, 0xBB3D, "CRC-16/ARC"}
	CRC16_AUG_CCITT   = Params{0x1021, 0x1D0F, false, false, 0x0000, 0xE5CC, "CRC-16/AUG-CCITT"}
	CRC16_BUYPASS     = Params{0x8005, 0x0000, false, false, 0x0000, 0xFEE8, "CRC-16/BUYPASS"}
	CRC16_CCITT_FALSE = Params{0x1021, 0xFFFF, false, false, 0x0000, 0x29B1, "CRC-16/CCITT-FALSE"}
	CRC16_CDMA2000    = Params{0xC867, 0xFFFF, false, false, 0x0000, 0x4C06, "CRC-16/CDMA2000"}
	CRC16_DDS_110     = Params{0x8005, 0x800D, false, false, 0x0000, 0x9ECF, "CRC-16/DDS-110"}
	CRC16_DECT_R      = Params{0x0589, 0x0000, false, false, 0x0001, 0x007E, "CRC-16/DECT-R"}
	CRC16_DECT_X      = Params{0x0589, 0x0000, false, false, 0x0000, 0x007F, "CRC-16/DECT-X"}
	CRC16_DNP         = Params{0x3D65, 0x0000, true, true, 0xFFFF, 0xEA82, "CRC-16/DNP"}
	CRC16_EN_13757    = Params{0x3D65, 0x0000, false, false, 0xFFFF, 0xC2B7, "CRC-16/EN-13757"}
	CRC16_GENIBUS     = Params{0x1021, 0xFFFF, false, false, 0xFFFF, 0xD64E, "CRC-16/GENIBUS"}
	CRC16_MAXIM       = Params{0x8005, 0x0000, true, true, 0xFFFF, 0x44C2, "CRC-16/MAXIM"}
	CRC16_MCRF4XX     = Params{0x1021, 0xFFFF, true, true, 0x0000, 0x6F91, "CRC-16/MCRF4XX"}
	CRC16_RIELLO      = Params{0x1021, 0xB2AA, true, true, 0x0000, 0x63D0, "CRC-16/RIELLO"}
	CRC16_T10_DIF     = Params{0x8BB7, 0x0000, false, false, 0x0000, 0xD0DB, "CRC-16/T10-DIF"}
	CRC16_TELEDISK    = Params{0xA097, 0x0000, false, false, 0x0000, 0x0FB3, "CRC-16/TELEDISK"}
	CRC16_TMS37157    = Params{0x1021, 0x89EC, true, true, 0x0000, 0x26B1, "CRC-16/TMS37157"}
	CRC16_USB         = Params{0x8005, 0xFFFF, true, true, 0xFFFF, 0xB4C8, "CRC-16/USB"}
	CRC16_CRC_A       = Params{0x1021, 0xC6C6, true, true, 0x0000, 0xBF05, "CRC-16/CRC-A"}
	CRC16_KERMIT      = Params{0x1021, 0x0000, true, true, 0x0000, 0x2189, "CRC-16/KERMIT"}
	CRC16_MODBUS      = Params{0x8005, 0xFFFF, true, true, 0x0000, 0x4B37, "CRC-16/MODBUS"}
	CRC16_X_25        = Params{0x1021, 0xFFFF, true, true, 0xFFFF, 0x906E, "CRC-16/X-25"}
	CRC16_XMODEM      = Params{0x1021, 0x0000, false, false, 0x0000, 0x31C3, "CRC-16/XMODEM"}
)

// Table is a 256-word table representing polinomial and algorithm settings for efficient processing.
type Table struct {
	params Params
	data   [256]uint16
}

func MakeTable(params Params) *Table {
	table := new(Table)
	table.params = params
	for n := 0; n < 256; n++ {
		crc := uint16(n) << 8
		for i := 0; i < 8; i++ {
			bit := (crc & 0x8000) != 0
			crc <<= 1
			if bit {
				crc ^= params.Poly
			}
		}
		table.data[n] = crc
	}
	return table
}

func MakeTable2222(params Params) *Table {
	table := new(Table)
	table.params = params
	n := 0
	for j := 0; j <= 31; j++ {
		for i := 8 * j; i < 8*j+8; i++ {
			table.data[n] = crchware(uint16(i))
			n++
		}
	}
	return table
}

func crchware(data uint16) uint16 {
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

// Init returns the initial value for CRC register corresponding to the specified algorithm.
func Init(table *Table) uint16 {
	return table.params.Init
}

// Update returns the result of adding the bytes in data to the crc.
func Update(crc uint16, data []byte, table *Table) uint16 {
	for _, d := range data {
		if table.params.RefIn {
			d = ReverseByte(d)
		}
		crc = crc<<8 ^ table.data[byte(crc>>8)^d]
	}
	return crc
}

func Update22222(accum uint16, arr []byte, table *Table) uint16 {
	accum = 0xffff
	for _, data := range arr {
		//accum = (accum << 8) ^ table.data[byte(accum>>8)^data]
		comb_val := (accum >> 8) ^ uint16(data)
		tmp := crchware(comb_val)
		accum = tmp ^ (accum << 8)
	}
	return accum
}

// Complete returns the result of CRC calculation and post-calculation processing of the crc.
func Complete(crc uint16, table *Table) uint16 {
	if table.params.RefOut {
		return ReverseUint16(crc) ^ table.params.XorOut
	}
	return crc ^ table.params.XorOut
}
func Complete222222(accum uint16, crctab *Table) uint16 {
	return accum
}

func Checksum(data []byte, table *Table) uint16 {
	crc := Init(table)
	crc = Update22222(crc, data, table)
	//return Complete(crc, table)
	return Complete222222(crc, table)
}
