package lapair

// source: http://latindictionary.wikidot.com/portable-dictionary

// Gender to array index:
//
//           | len 3 | len 2 | len 1
// ---------------------------------
// Masculine | 0     | 0     | 0
// Feminine  | 1     | 0     | 0
// Neuter    | 2     | 1     | 0

var latinA = [][]string{
	{"absurdus", "absurda", "absurdum"},       // absurd, silly, pointless
	{"acer", "acris"},                         // fierce, sharp, irritating, pungent, bitter
	{"aequus", "aequa", "aequum"},             // equal, fair, level
	{"aeternus", "aeterna", "aeternum"},       // eternal, immortal
	{"alius", "alia", "aliud"},                // another, other
	{"alter", "altera", "alterum"},            // the other (of two); second
	{"anglicus", "anglica", "anglicum"},       // english
	{"antiquus", "antiqua", "antiquum"},       // ancient, old
	{"anxius", "anxia", "anxium"},             // anxious
	{"apertus", "aperta", "apertum"},          // open
	{"aridus", "arida", "aridum"},             // dry
	{"artus", "arta", "artum"},                // narrow, close, tight
	{"ater", "atra", "atrum"},                 // black, dark
	{"attonitus", "attonita", "attonitum"},    // astonished, awestruck
	{"audax", "audacis"},                      // brave, bold, daring
	{"avarus", "avara", "avarum"},             // greedy
	{"beatus", "beata", "beatum"},             // happy, fortunate, prosperous
	{"bellus", "bella", "bellum"},             // pleasant, pretty, handsome
	{"benignus", "benigna", "benignum"},       // kind
	{"bonus", "bona", "bonum"},                // good
	{"brevis", "breve"},                       // short, brief
	{"caecus", "caeca", "caecum"},             // blind; secret, dark
	{"candidus", "candida", "candidum"},       // shining, bright, white, beautiful
	{"carus", "cara", "carum"},                // dear, expensive
	{"celer", "celeris"},                      // fast, quick, swift, speedy
	{"certus", "certa", "certum"},             // certain
	{"clarus", "clara", "clarum"},             // clear, bright, shining, famous
	{"contentus", "contenta", "contentum"},    // content, comfortable
	{"crudelis", "crudele"},                   // cruel
	{"curvus", "curva", "curvum"},             // bent, curved
	{"defessus", "defessa", "defessum"},       // exhausted, tired
	{"desertus", "deserta", "desertum"},       // deserted
	{"dexter", "dextra", "dextrum"},           // right (hand); skillful
	{"difficilis", "difficile"},               // difficult
	{"dignus", "digna", "dignum"},             // worthy, deserving
	{"discolor", "discoloris"},                // of a different color
	{"dives", "divitis"},                      // rich
	{"divinus", "divina", "divinum"},          // divine
	{"dubius", "dubia", "dubium"},             // doubtful, unsure
	{"dulcis", "dulce"},                       // sweet
	{"ebrius", "ebrium", "ebria"},             // drunk, intoxicated
	{"fabulosus", "fabulosa", "fabulosum"},    // fabulous (from a story)
	{"facilis", "facile"},                     // easy
	{"falsus", "falsa", "falsum"},             // false
	{"fatuus", "fatua", "fatuum"},             // foolish, stupid
	{"felix", "felicis"},                      // happy, fortunate
	{"ferox", "ferocis"},                      // fierce, wild, savage
	{"fessus", "fessa", "fessum"},             // tired, exausted
	{"fidelis", "fidele"},                     // faithful, loyal
	{"fortis", "forte"},                       // brave, strong, powerful, courageous
	{"fractus", "fracta", "fractum"},          // broken
	{"globosus", "globosa", "globosum"},       // spherical
	{"graecus", "graeca", "graecum"},          // greek
	{"gratus", "grata", "gratum"},             // pleasing, grateful
	{"gravis", "grave"},                       // heavy, serious
	{"hibernicus", "hibernica", "hibernicum"}, // irish
	//{"hic", "haec", "hoc"},                          // this
	{"hilaris", "hilaris"},                    // cheerful
	{"honestus", "honesta", "honestum"},       // honest, honorable
	{"horribilis", "horribilis", "horribile"}, // horrible, terrible
	{"humanus", "humana", "humanum"},          // kind, humane; human
	{"iaponicus", "iaponica", "iaponicum"},    // japanese
	{"ignarus", "ignara", "ignarum"},          // ignorant, unaware
	{"ignavus", "ignava", "ignavum"},          // lazy
	{"immemor", "immemoris"},                  // forgetful of
	{"imprudens", "imprudentis"},              // imprudent, shameless
	{"incolumis", "incolumis"},                // safe
	{"infelix", "infelicis"},                  // unhappy, unlucky, unfortunate
	{"ingens", "ingentis"},                    // huge
	{"ingratus", "ingrata", "ingratum"},       // thankless, ungrateful
	{"invitus", "invita", "invitum"},          // unwilling
	{"iocosus", "iocosa", "iocosum"},          // humorous, playful
	{"iratus", "irata", "iaratum"},            // angry
	{"iucundus", "iucunda", "iucundum"},       // pleasant, agreeable, delightful
	{"laboriosus", "laboriosa", "laboriosum"}, // laborious, toilsome
	{"laetus", "laeta", "laetum"},             // happy, joyful
	{"latus", "lata", "latum"},                // wide, broad
	{"lepidus", "lepida", "lepidum"},          // pleasant, elegant, witty, neat
	{"levis", "levis"},                        // light, slight, easy, trivial
	{"liber", "libera", "liberum"},            // free
	{"magnus", "magna", "magnum"},             // large, great, important
	{"maior", "maioris"},                      // bigger, greater
	{"malus", "mala", "malum"},                // bad, evil
	{"maximus", "maxima", "maximum"},          // greatest, biggest
	{"medius", "media", "medium"},             // middle, in the middle of
	{"melior", "melioris"},                    // better
	{"merus", "mera", "merum"},                // pure, undiluted
	//{"meus", "mea", "meum"},                         // my
	{"minimus", "minima", "minimum"},    // smallest
	{"minor", "minoris"},                // smaller
	{"mirus", "mira", "mirum"},          // wonderful, marvelous
	{"misellus", "misella", "misellum"}, // poor little
	{"miser", "misera", "miserum"},      // sad, unhappy, miserable, wretched
	{"mortalis", "mortalis"},            // mortal
	{"mortuus", "mortua", "mortuum"},    // mortal
	//{"multus", "multa", "multum"},                   // much, many
	{"necesse"}, // necessary
	//{"neuter", "neutra", "neutrum"},                 // neither
	{"niger", "nigra", "nigrum"},       // black, dark
	{"nobilis", "nobilis"},             // noble
	{"nonulli", "nonullae", "nonulla"}, // some
	{"noster", "nostra", "nostrum"},    // ours
	{"notus", "nota", "notum"},         // known, famous
	{"novus", "nova", "novum"},         // new, strange
	//{"nullus", "nulla", "nullum"},                   // none, no
	//{"nuptialis", "nuptialis"},                      // nuptial, of a wedding
	{"obstinatus", "obstinata", "obstinatum"}, // firm, resolute; stubborn
	{"occupatus", "occupata", "occupatum"},    // occupied, busy
	{"omnipotens", "omnipotentis"},            // omnipotent
	{"omnis", "omnis"},                        // everything, all
	{"optimus", "optima", "optimum"},          // best
	{"otiosus", "otiosa", "otiosum"},          // at leisure, free
	//{"par", "paris"},                                // equal, like
	{"paratus", "parata", "paratum"}, // prepared, ready
	{"parvus", "parva", "parvum"},    // small, little
	//{"pauci", "paucae", "pauca"},                    // few
	{"pauper", "pauperis"},                 // poor
	{"peior", "peioris"},                   // worse
	{"perpetuus", "perpetua", "perpetuum"}, // perpetual, long lasting, continuous, uninterrupted
	{"pervicax", "pervicacis"},             // obstinate, wilful, stubborn
	{"pessimus", "pessima", "pessimum"},    // worst
	{"pius", "pia", "pium"},                // dutiful, conscientious
	{"plenus", "plena", "plenum"},          // full, abundant, generous
	//{"plerique", "pleraeque", "pleraque"},           // several
	//{"plurimus", "plurima", "plurimum"},             // most
	//{"plus", "ploris"},                              // more
	{"potens", "potentis"}, // powerful, potent
	//{"potis"},                                       // possible, able
	{"prior", "prioris"},                      // prior, former, previous
	{"propinquus", "propinqua", "propinquum"}, // relative, kindred
	{"pudicus", "pudica", "pudicum"},          // pure, chaste, modest
	{"pulcher", "pulchra", "pulchrum"},        // handsome, beautiful, fine
	{"purus", "pura", "purum"},                // pure
	//{"qualis", "qualis"},                            // what sort of
	//{"quantus", "quanta", "quantum"},                // how great? how large?
	//{"quot"},                                        // how many; as many as
	{"reus", "rea", "reum"},                // accused
	{"rigidus", "rigida", "rigidum"},       // stiff, rigid
	{"romanus", "romana", "romanum"},       // roman, latin
	{"salvus", "salva", "salvum"},          // safe, sound
	{"sanus", "sana", "sanum"},             // healthy, sane, sound
	{"sapiens", "sapientis"},               // wise, sensible, judicious
	{"saucius", "saucia", "saucium"},       // hurt, ill
	{"scelestus", "scelesta", "scelestum"}, // wicked
	{"severus", "severa", "severum"},       // severe, serious
	{"sinister", "sinistra", "sinistrum"},  // left
	{"solus", "sola", "solum"},             // alone, only
	{"spatiosus", "spatiosa", "spatiosum"}, // roomy, ample, large
	{"stultus", "stulta", "stultum"},       // foolish
	{"summus", "summa", "summum"},          // greatest
	{"superbus", "superba", "superbum"},    // proud
	//{"suus", "sua", "suum"},                         // his, her, its, their
	//{"talis", "talis"},                              // such
	{"tenebricosus", "tenebricosa", "tenebricosum"}, // dark, gloomy
	{"tenuis", "tenuis"},                            // thin, fine; small
	{"territus", "territa", "territum"},             // terrified, scared
	{"timidus", "timida", "timidum"},                // timid
	//{"tot"},                                         // so many, as many
	//{"totus", "tota", "totum"},                      // whole, entire
	{"tristis", "tristis"}, // sad, gloomy
	//{"tuus", "tua", "tuum"},                         // your (singular)
	//{"ullus", "ulla", "ullum"},                      // any
	{"ultimus", "ultima", "ultimum"},       // last, final
	{"universus", "universa", "universum"}, // all
	//{"uter", "utra", "utrum"},                       // which of two, either
	//{"varius", "varia", "varium"},                   // various, diverse, different
	{"vehemens", "vehementis"},          // fierce, energetic
	{"venustus", "venusta", "venustum"}, // charming, attractive beautiful
	{"verus", "vera", "verum"},          // real, true
	//{"vester", "vestra", "vestrum"},                 // your (plural)
	{"vilis", "vilis"},         // cheap, worthless, poor, common
	{"virilis", "virilis"},     // manly, of man
	{"vivus", "viva", "vivum"}, // alive, living
}
