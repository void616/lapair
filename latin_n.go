package lapair

// source: http://latindictionary.wikidot.com/portable-dictionary

// m = Masculine
// f = Feminine
// n = Neuter

var latinN = []struct {
	v string
	g Gender
}{
	{"adulescentia", Feminine}, // youth, youthfulness
	{"aedificium", Neuter},     // building
	{"aestas", Feminine},       // summer
	{"aestus", Masculine},      // heat
	{"aetas", Feminine},        // age, lifetime
	{"aevum", Neuter},          // age, lifetime; eternity
	{"ager", Masculine},        // field, farm
	{"agora", Feminine},        // agora, city center
	{"agricola", Masculine},    // farmer
	{"ala", Feminine},          // wing
	{"amica", Feminine},        // girlfriend, female friend
	{"amicitia", Feminine},     // friendship
	{"amicus", Masculine},      // friend
	{"amor", Masculine},        // love
	//{"anglia", Feminine},       // england
	{"animus", Masculine},      // soul, spirit, mind; (plural) high spirits, pride, courage
	{"annus", Masculine},       // year
	{"anxietas", Feminine},     // anxiety
	{"apis", Feminine},         // bee
	{"aqua", Feminine},         // water
	{"aranea", Feminine},       // spider
	{"arbor", Feminine},        // tree
	{"arcuballista", Feminine}, // crossbow
	{"ardor", Masculine},       // heat, brightness; passion (figurative)
	{"argentum", Neuter},       // silver, money
	{"arma", Neuter},           // arms, weapons
	{"auditor", Masculine},     // listener, audience
	{"aula", Feminine},         // courtyard
	{"aura", Feminine},         // breeze
	{"auxilium", Neuter},       // help
	{"avis", Feminine},         // bird
	{"avus", Masculine},        // grandfather
	{"balnea", Neuter},         // baths
	{"basium", Neuter},         // kiss (usually loud)
	{"bellum", Neuter},         // war
	{"bona", Neuter},           // goods, property
	{"bos", Feminine},          // cow
	{"caelum", Neuter},         // sky, heaven
	{"calceus", Masculine},     // shoe
	{"campus", Masculine},      // field, plane
	{"canis", Masculine},       // dog
	{"capillus", Masculine},    // hair
	{"caput", Neuter},          // head, chapter
	{"carcer", Masculine},      // prison, barrier; starting area for a race
	{"carmen", Neuter},         // song, poem
	{"casa", Feminine},         // house, cottage
	{"castra", Neuter},         // camp
	{"casus", Masculine},       // fall, case, event, misfortune
	{"caupona", Feminine},      // inn
	{"causa", Feminine},        // cause, reason
	{"cena", Feminine},         // dinner
	{"cerva", Feminine},        // deer
	{"cervix", Feminine},       // neck
	{"cibus", Masculine},       // food
	{"civis", Feminine},        // citizen
	{"collis", Masculine},      // hill
	{"collum", Neuter},         // neck
	{"colonus", Masculine},     // farmer
	{"comes", Feminine},        // friend, companion
	{"consilium", Neuter},      // plan, purpose, counsel, advice, judgement, wisdom
	{"copia", Feminine},        // abundance, supply; troops, forces
	{"cor", Neuter},            // heart
	{"cornu", Neuter},          // horn, army flank
	{"corona", Feminine},       // crown, garland
	{"corpus", Neuter},         // body
	{"crimen", Neuter},         // crime
	{"crus", Neuter},           // leg, shin
	{"culina", Feminine},       // kitchen
	{"culpa", Feminine},        // fault, blame
	{"cupido", Feminine},       // desire, longing, passion; cupid
	{"cura", Feminine},         // care, attention, caution, anxiety
	{"currus", Masculine},      // chariot
	{"cursus", Masculine},      // running, track, course
	{"custos", Masculine},      // guard
	{"dea", Feminine},          // goddess
	{"defensor", Masculine},    // defender, protector
	{"deliciae", Feminine},     // delight, pleasure; sweetheart
	{"delubrum", Neuter},       // temple, sanctuary, shrine
	{"deus", Masculine},        // god
	{"dextra", Feminine},       // right hand, right hand side
	{"dictator", Masculine},    // dictator
	{"dies", Feminine},         // day
	{"diligentia", Feminine},   // diligence, care
	{"discipulus", Masculine},  // student
	{"divitiae", Feminine},     // riches, wealth
	{"dolor", Masculine},       // pain, grief
	{"domina", Feminine},       // mistress
	{"domus", Feminine},        // house, home
	{"donum", Neuter},          // gift, present
	{"dux", Masculine},         // leader
	{"eloquentia", Feminine},   // eloquence
	{"epistola", Feminine},     // letter, epistle
	{"eques", Masculine},       // knight, horseman
	{"equus", Masculine},       // horse
	{"error", Masculine},       // errand, error
	{"exemplum", Neuter},       // example, sample, model
	{"exitium", Neuter},        // destruction, ruin
	{"fabula", Feminine},       // story, tale, play
	{"facilitas", Feminine},    // ease, facility, easiness
	{"fama", Feminine},         // fame, reputation
	{"fames", Feminine},        // hunger
	{"familia", Feminine},      // family
	{"fas", Neuter},            // (divine) law
	{"fatum", Neuter},          // fate, death
	{"femina", Feminine},       // woman
	{"ferrum", Neuter},         // iron, sword
	{"fides", Feminine},        // trust, faith, belief
	{"filia", Feminine},        // daughter
	{"filius", Masculine},      // son
	{"finis", Masculine},       // end, limit, border; territory (plural)
	{"flamma", Feminine},       // flame, fire, torch, star
	{"flos", Masculine},        // flower
	{"flumen", Neuter},         // river
	{"foedus", Neuter},         // treaty, pact
	{"fons", Masculine},        // fountain, spring
	{"forma", Feminine},        // form, shape, beauty
	{"fortuna", Feminine},      // fortune, luck
	{"forum", Neuter},          // market place, public center
	{"frater", Masculine},      // brother
	{"frumentum", Neuter},      // grain
	{"funus", Neuter},          // funeral
	{"furor", Masculine},       // madness, insanity
	{"gaudium", Neuter},        // joy
	{"genu", Neuter},           // knee
	{"genus", Neuter},          // gender, descendancy, kind, sort, class
	{"gladius", Masculine},     // sword
	{"gloria", Feminine},       // glory, fame
	{"gremium", Neuter},        // bosom, lap
	{"hasta", Feminine},        // spear
	//{"hibernia", Feminine},     // ireland
	{"hiems", Feminine},     // winter
	{"hobbitus", Masculine}, // hobbit
	{"homo", Masculine},     // man, mankind, human
	{"honor", Masculine},    // honor, esteem, public office
	{"hortus", Masculine},   // garden
	{"hostis", Masculine},   // enemy
	{"humus", Feminine},     // ground
	{"ianua", Feminine},     // door
	//{"iaponia", Feminine},      // japan
	{"ignis", Masculine},      // fire; lightning
	{"imperium", Neuter},      // order, command; power
	{"incendium", Neuter},     // fire
	{"infinitum", Neuter},     // infinity
	{"ingenium", Neuter},      // nature, innate talent
	{"inimicus", Masculine},   // enemy
	{"iniuria", Feminine},     // injury, offense, injustice, wrongdoing
	{"insidiae", Feminine},    // ambush, plot, treachery
	{"intestina", Neuter},     // intestines, guts, entrails
	{"iocus", Masculine},      // joke, jest
	{"ira", Feminine},         // ire, anger
	{"iter", Neuter},          // journey
	{"iudicium", Neuter},      // trial, court; judgement, opinion, sentence
	{"iugulum", Neuter},       // throat
	{"iuppiter", Masculine},   // jupiter
	{"ius", Neuter},           // law, right, justice
	{"iuvenis", Masculine},    // youth, young man
	{"lac", Neuter},           // milk
	{"lacrima", Feminine},     // tear
	{"laetitia", Feminine},    // happiness, joy, delight, elation
	{"lapis", Masculine},      // stone
	{"laus", Feminine},        // praise, glory, fame
	{"lectus", Masculine},     // bed, couch
	{"legatus", Masculine},    // deputy, officer, envoy, diplomat
	{"legio", Feminine},       // legion
	{"leo", Masculine},        // lion
	{"lex", Feminine},         // law
	{"libellus", Masculine},   // small book, notebook; diary
	{"liber", Masculine},      // book
	{"libertas", Feminine},    // liberty, freedom
	{"libertus", Masculine},   // freedman
	{"limen", Neuter},         // threshold, doorway, limit
	{"linea", Feminine},       // line
	{"lingua", Feminine},      // tongue; language
	{"littera", Feminine},     // letter (alphabet)
	{"litterae", Feminine},    // letter, epistle; literature
	{"litterator", Masculine}, // schoolmaster, grammarian, teacher
	{"litus", Neuter},         // shore, coast
	{"locus", Masculine},      // place, location
	{"ludus", Masculine},      // school, game
	{"lumen", Neuter},         // light
	{"lux", Feminine},         // light
	{"magister", Masculine},   // teacher
	{"mala", Neuter},          // troubles, evils
	{"manus", Feminine},       // hand
	{"maritus", Masculine},    // husband
	{"mater", Feminine},       // mother
	{"matrona", Feminine},     // married woman
	{"membrum", Neuter},       // limb; portion; room; member
	{"mens", Feminine},        // mind
	{"mensa", Feminine},       // table; meal
	{"mensis", Masculine},     // month
	{"meridies", Masculine},   // midday
	{"meta", Feminine},        // turn; a pillar in a circus race
	{"miles", Masculine},      // soldier
	{"mille", Feminine},       // mile
	{"moenia", Neuter},        // walls of a city
	{"mons", Masculine},       // mountain
	{"monumentum", Neuter},    // monument
	{"mora", Feminine},        // delay
	{"mors", Feminine},        // death
	{"morsus", Masculine},     // bite, grip
	{"mos", Masculine},        // habit, custom, manner
	{"mulier", Feminine},      // woman
	{"mundus", Masculine},     // world
	{"munus", Neuter},         // gift; gladiatorial show
	{"murus", Masculine},      // wall
	{"nasus", Masculine},      // nose
	{"natio", Feminine},       // birth; nation, people
	{"natus", Masculine},      // son
	{"nauta", Masculine},      // sailor
	{"navis", Feminine},       // ship
	{"negotium", Neuter},      // business, affair
	{"nemo", Feminine},        // nobody, no one
	{"nemus", Neuter},         // forest
	//{"nihil", "none"},         // nothing
	{"nomen", Neuter},         // name, title
	{"nox", Feminine},         // night
	{"nubes", Feminine},       // cloud
	{"numerus", Masculine},    // number
	{"nundinae", Feminine},    // market day
	{"nuntius", Masculine},    // messenger
	{"obses", Feminine},       // hostage
	{"oculus", Masculine},     // eye
	{"officium", Neuter},      // duty, service
	{"ops", Masculine},        // power or ability to help; resources (plural)
	{"opus", Neuter},          // work
	{"oratio", Feminine},      // speech, lecture
	{"orator", Masculine},     // orator, speaker
	{"ordo", Masculine},       // order, rank, line
	{"os", Neuter},            // mouth
	{"os", Neuter},            // bone
	{"otium", Neuter},         // leisure, peace
	{"paedagogus", Masculine}, // tutor
	{"pars", Feminine},        // part, portion, share
	{"passer", Masculine},     // sparrow, bird
	{"pater", Masculine},      // father
	{"patria", Feminine},      // fatherland, native land, country
	{"pectus", Neuter},        // chest, breast; heart, mind (fig)
	{"pecunia", Feminine},     // money
	{"periculum", Neuter},     // danger, risk
	{"pes", Masculine},        // foot
	{"pestis", Feminine},      // plague, pest; ruin, destruction
	{"philosophia", Feminine}, // philosophy
	{"pietas", Feminine},      // piety, sense of duty; filial love
	{"poena", Feminine},       // punishment
	{"poeta", Masculine},      // poet
	{"pons", Masculine},       // bridge
	{"populus", Masculine},    // people, nation
	{"porta", Feminine},       // gate, entrance
	{"postis", Masculine},     // door, post, doorpost
	{"praesidium", Neuter},    // garrison
	{"praetor", Masculine},    // praetor (elected magistrate)
	{"provincia", Feminine},   // duty, province
	{"pudor", Masculine},      // shame, modesty
	{"puella", Feminine},      // girl
	{"puer", Masculine},       // boy
	{"pugna", Feminine},       // a fight, battle, dispute
	{"pumex", Masculine},      // pumice, stone, rock
	{"puppis", Feminine},      // ship, stern
	{"quies", Feminine},       // rest, peace, quiet
	{"ratio", Feminine},       // reckoning, account, reason, judgement, consideration, system, manner, method
	{"regia", Feminine},       // palace
	{"regina", Feminine},      // queen
	{"remedium", Neuter},      // cure, remedy
	{"res", Feminine},         // object, thing; matter, property
	{"rex", Masculine},        // king
	{"risus", Masculine},      // smile, laugh
	{"roma", Feminine},        // rome
	{"rosa", Feminine},        // rose
	{"rumor", Masculine},      // rumor, report
	{"sacrificium", Neuter},   // sacrifice
	{"saeculum", Neuter},      // generation, lifetime, age
	{"salus", Feminine},       // safety, health, greeting
	{"sanguis", Masculine},    // blood
	{"sapientia", Feminine},   // wisdom
	{"saxum", Neuter},         // rock, boulder
	{"scacci", Masculine},     // chess
	{"scelus", Neuter},        // crime
	{"scriptor", Masculine},   // writer, author
	{"senator", Masculine},    // senator
	{"senatus", Masculine},    // the senate
	{"senex", Masculine},      // old man
	{"sententia", Feminine},   // feeling, thought, opinion, vote, sentence
	{"sepulchrum", Neuter},    // tomb
	{"serpens", Masculine},    // serpent, snake, dragon
	{"serva", Feminine},       // slavewoman
	{"servus", Masculine},     // slave
	{"signum", Neuter},        // sign, signal, zeal
	{"silva", Feminine},       // forest, woods
	{"sinistra", Feminine},    // left hand
	{"sol", Masculine},        // sun
	{"somnium", Neuter},       // dream
	{"soror", Feminine},       // sister
	{"species", Feminine},     // appearance
	{"spes", Feminine},        // hope
	{"spiritus", Masculine},   // breath, breathing; breeze, air; inspiration; character; arrogance
	{"statua", Feminine},      // statue
	{"stella", Feminine},      // star; constellation
	{"studium", Neuter},       // study, pursuit
	{"tabula", Feminine},      // tablet
	{"taurus", Masculine},     // bull
	{"tempestas", Feminine},   // storm
	{"templum", Neuter},       // temple
	{"tempus", Neuter},        // time
	{"tenebrae", Feminine},    // darkness, night, shadows; death, blindness
	{"tergum", Neuter},        // back, rear
	{"timor", Masculine},      // fear
	{"toga", Feminine},        // toga
	{"tonitrus", Masculine},   // thunder
	{"tunica", Feminine},      // tunic
	{"tyrannus", Masculine},   // tyrant
	{"umbra", Feminine},       // shade, shadow, spirit
	{"uxor", Feminine},        // wife
	{"venenum", Neuter},       // drug, poison
	{"venia", Feminine},       // pardon
	{"ventus", Masculine},     // wind
	{"venus", Feminine},       // charm, beauty; love, mating; goddess of love
	{"venustas", Feminine},    // charm, beauty
	{"ver", Neuter},           // spring (season)
	{"verber", Neuter},        // whip, lash
	{"verbum", Neuter},        // word
	{"veritas", Feminine},     // truth
	{"viator", Feminine},      // traveler
	{"vicinus", Masculine},    // neighbor
	{"victoria", Feminine},    // victory
	{"villa", Feminine},       // villa, house
	{"vir", Masculine},        // man, hero
	{"virgo", Feminine},       // girl, virgin
	{"virtus", Feminine},      // manliness, courage, excellence, character, virtue
	{"vita", Feminine},        // life
	{"vitium", Neuter},        // fault, crime, vice
	{"voluptas", Feminine},    // pleasure, enjoyment; entertainment (plural)
	{"votum", Neuter},         // vow, prayer; longing
	{"vox", Feminine},         // voice
	{"vulgus", Neuter},        // crowd, people
	{"vulnus", Neuter},        // wound
	{"vultus", Masculine},     // face, expression
}
