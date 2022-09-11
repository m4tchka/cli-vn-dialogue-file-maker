package main

// type chObj struct {}
type SceneObj struct {
	Id    int           `json:"id"`
	Scene []DialogueObj `json:"scene"`
}
type DialogueObj struct {
	// Only Name & Dialogue fields mandatory, rest will remain as empty strings/ slice until conversion to JSON upon where they will be removed.
	Name       string `json:"Name"`
	Dialogue   string `json:"Dialogue"`
	Background string `json:"Background,omitempty"`
	Question   string `json:"Question,omitempty"`
	/*
		NOTE: In the actual VN, if there is a question key then
		WARNING: THE NAME AND DIALOGUE KEYS ARE NOT RENDERED
		So if a particular dialogueObj is expected to be a question, then any name and dialogue can be put in the field,
		since it won't be displayed anyway. WARNING:
	*/
	Options []OptionObj `json:"Options,omitempty"`
	//Sprites SpriteObj
}
type OptionObj struct {
	// All fields mandatory
	Text       string `json:"Text"`
	Next       int    `json:"Next"`
	LuckChange int    `json:"LuckChange"`
	MinLuck    int    `json:"MinLuck"`
}
