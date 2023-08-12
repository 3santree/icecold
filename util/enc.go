package icecold

import "golang.org/x/exp/slices"

var wordlist = []string{"acllivirt_79", "praigo", "gabsalyam777", "leomascarenhas", "Eriamjh", "elderwulf", "russw79", "fearfio", "sergix", "IDTia", "stroiiiim", "krofty16", "darksceen", "PknfId", "refuma59", "saxa8097", "clhommes", "froa0jrup8y", "kevocfe337", "nafzal3", "liherman", "sv.valerievna", "rxmao11", "sidddv", "lavanda", "ahmetcayir", "elxordi", "aganykin", "dnhkbgk", "randyj1022", "craig020", "0IVIV0P1oOgl", "frekney", "micknonslamic", "antikozel", "danny53", "kapito02", "erikrune", "jarev", "griffypop", "maidaneh", "aristov.andr", "zr7willtree", "thasman", "buby-1", "onlyyou-al", "nathair1", "liliya.art08", "dorosch-vadim", "gvxfvm", "dosia_42", "elviranuralieva", "urforuo", "bruno.macedo", "ho4upiva", "martymm2", "chalik", "jwoolfe2000", "paragonpc", "hairy.womble", "gizemakar9", "hdfxs", "mongasumit", "bla3e.str", "chereago", "marko_mark", "ww_ww14", "fx33i", "bfaunt", "vzj82zra", "asibert3", "pavelsev", "j.harker2", "bigulit", "john_bk", "bbndlihelzahmpz", "anklbiters3", "bwpowell44", "Jena113", "kienergy", "edemers", "aldar675", "bigbigcar", "cloons", "brucebarr", "gadgadter", "director.kanavto", "soad1623", "ovl.of.death", "eliselutz375", "bmroberts7", "malarkey21", "wfb0421", "habi_rusta", "nick_murdock", "ethota", "psm3234", "schnebba", "rachevvladimir", "tvk100", "codecrackers", "dv_com", "3969427", "irondimas", "jurasrowerek", "crypokids", "unit011", "vmpnarayan", "sirpussyalot", "millervalera", "AefftKidsz", "ready_ryan", "zurik_mikava", "garmoniydushi", "qqyoitzlaeqq", "pnitram2696079", "janellen83", "tammy30", "nelsonjr2003", "molcal520", "svv-vlad", "kosmos.step", "thais_gc7", "HnyepB", "Guido42", "tu_an1610", "fmodovuz", "mounirgaabaji", "eal001", "dcbaabcd_5", "jc.lhof", "lenaandreeva1977", "dmh7868", "claudia_schifferd", "rykina53", "scottnlins1", "chaiwat3555", "gul_miral", "emerimike", "moodaskeezer", "lev-stv", "cenuit", "tecnoriarreda", "norbie1", "paulo_bonitao_18", "karina_usypova", "allbig1021", "BERTO42", "romulks", "saxoner", "gionica", "alexeygoncharov", "ckjen21", "arte-300", "mironov.arnur", "jakob759", "amnesia-k", "siberixe", "rekhac1988", "blackdog42", "greeshan", "fntudu.oimfnu", "Noahline", "pkelley7", "dbcbas", "ivycat", "rasyidahruslee", "babaxyanc_kristin", "nidomamita", "rowdyartist", "skashnikov", "Kalininromv", "SCMSCM", "fuzypikl", "hairyot", "GENE09", "pikechas", "fourever", "70003865", "jrickard", "slava6604", "jsc2000", "rick.holmes", "jhnngee", "Pedimiki", "amesan22", "GCL6D23W", "silver-fly", "linxinzztv12", "henry65", "nakedcity", "bandits21", "ab69836068", "TygReTgeR", "mimi69", "Citizen-inica", "buster__21", "Miriamguizar20", "lopez_torres55", "epityxia", "huskey1", "harithaw", "genja-887", "45hzwg", "frizzlef", "paulexnder", "b.c.n.g.edge", "rachelas", "ch179k", "goodheavens", "minduro69", "08263407", "Coolyen", "tw12345", "yuliya794", "trefoils", "xadgeevyura", "chebbi8076", "denzelday", "Johan87", "tairo", "cstgroupsrl", "gujied", "acesayshi", "wilma396", "1spec912", "conya123123", "rogerhcab", "sumbiloncarmela", "GIGILM", "dhanavut", "laouki", "shlemazlizrailevich", "wendell231", "Leontjev.roma2010", "psittakossss", "demkina.l", "dadits1985", "brianwright", "wkdennis", "tonyco68", "victor16", "inge.mangel", "lewis_grimley-2k8", "the_omaha_scientist", "slasher21", "hugo_almeida", "tukaevanton", "DemonYang", "dcf1966", "titerillo", "nmazuRok", "tadeyko14", "cugqatulht", "altwixer", "notoy"}

func Enc(sc []byte) []string {
	enclist := make([]string, len(sc))
	for i := 0; i < len(sc); i++ {
		enclist[i] = wordlist[int(sc[i])]
	}
	return enclist
}

func Dec(el []string) []byte {
	sc := make([]byte, len(el))
	for i := 0; i < len(el); i++ {
		sc[i] = byte(slices.Index(wordlist, el[i]))
	}
	return sc
}
