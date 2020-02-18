package laofont

// Lao-ASCII mapping with Unicode table
var asciiTable = map[rune]string{
	'¡':      "ກ",
	'¢':      "ຂ",
	'£':      "ຄ",
	'¤':      "ງ",
	'¥':      "ຈ",
	'¦':      "ສ",
	'§':      "ຊ",
	'¨':      "ຍ",
	'©':      "ດ",
	'ª':      "ຕ",
	'«':      "ຖ",
	'ê':      "ທ",
	'ß':      "ນ", // ນ with newer ascii
	'\u00ad': "ນ", // ນ with older ascii
	'®':      "ບ",
	'¯':      "ປ",
	'°':      "ຜ",
	'±':      "ຝ",
	'²':      "ພ",
	'³':      "ຟ",
	'´':      "ມ",
	'µ':      "ຢ",
	'ì':      "ລ",
	'ë':      "ຣ",
	'¸':      "ວ",
	'¹':      "ຫ",
	'º':      "ອ",
	'»':      "ຮ",
	'½':      "ະ",
	'¾':      "າ",
	'ò':      "ິ",
	'ó':      "ີ",
	'ô':      "ຶ",
	'õ':      "ື",
	'÷':      "ຸ",
	'ø':      "ູ",
	'À':      "ເ",
	'Á':      "ແ",
	'Â':      "ໂ",
	'ð':      "ໍ",
	'ñ':      "ັ",
	'ö':      "ົ",
	'Ã':      "ໃ",
	'Ä':      "ໄ",
	'ú':      "່",
	'È':      "່",
	'û':      "້",
	'É':      "້",
	'ü':      "໊",
	'ý':      "໋",
	'Î':      "ໜ",
	'Ï':      "ໝ",
	'Í':      "ຫຼ",
	'ù':      "ຼ",
	'š':      "ີ້",
	'™':      "ິ້",
	'¿':      "ຳ",
	'‡':      "ຶ່",
	'Ò':      "ໍ່",
	'¼':      "ຽ",
	'˜':      "ັ້",
}

func asciiToUnicode(x rune) string {
	//  check if ASCII x not present in table then return x else return y from mapping table.
	y, ok := asciiTable[x]
	if !ok {
		return string(x)
	}
	return y
}

// ASCIIToUnicode ...
func ASCIIToUnicode(x string) string {
	var y string
	for _, r := range x {
		y += asciiToUnicode(r)
	}
	return y
}
