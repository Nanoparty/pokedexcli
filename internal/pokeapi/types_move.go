package pokeapi

type Move struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Accuracy      int    `json:"accuracy"`
	EffectChance  any    `json:"effect_chance"`
	Pp            int    `json:"pp"`
	Priority      int    `json:"priority"`
	Power         int    `json:"power"`
	ContestCombos struct {
		Normal struct {
			UseBefore []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"use_before"`
			UseAfter any `json:"use_after"`
		} `json:"normal"`
		Super struct {
			UseBefore any `json:"use_before"`
			UseAfter  any `json:"use_after"`
		} `json:"super"`
	} `json:"contest_combos"`
	ContestType struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"contest_type"`
	ContestEffect struct {
		URL string `json:"url"`
	} `json:"contest_effect"`
	DamageClass struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"damage_class"`
	EffectEntries []struct {
		Effect      string `json:"effect"`
		ShortEffect string `json:"short_effect"`
		Language    struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"effect_entries"`
	EffectChanges []any `json:"effect_changes"`
	Generation    struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"generation"`
	Meta struct {
		Ailment struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ailment"`
		Category struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"category"`
		MinHits       any `json:"min_hits"`
		MaxHits       any `json:"max_hits"`
		MinTurns      any `json:"min_turns"`
		MaxTurns      any `json:"max_turns"`
		Drain         int `json:"drain"`
		Healing       int `json:"healing"`
		CritRate      int `json:"crit_rate"`
		AilmentChance int `json:"ailment_chance"`
		FlinchChance  int `json:"flinch_chance"`
		StatChance    int `json:"stat_chance"`
	} `json:"meta"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PastValues         []any `json:"past_values"`
	StatChanges        []any `json:"stat_changes"`
	SuperContestEffect struct {
		URL string `json:"url"`
	} `json:"super_contest_effect"`
	Target struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"target"`
	Type struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"type"`
	LearnedByPokemon []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"learned_by_pokemon"`
	FlavorTextEntries []struct {
		FlavorText string `json:"flavor_text"`
		Language   struct {
			URL  string `json:"url"`
			Name string `json:"name"`
		} `json:"language"`
		VersionGroup struct {
			URL  string `json:"url"`
			Name string `json:"name"`
		} `json:"version_group"`
	} `json:"flavor_text_entries"`
}