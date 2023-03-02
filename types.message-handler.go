package mom

// MessageHandler defines the function layout for a message processor
// returning an object will return the JSON equivalent of that object back to the requesting service, if it is still listening
// returning an error will return the error back to the requesting service, if it is still listening
// if a panic is encountered system will automatically just return it to avoid service crashes
// if you want to exit service instead use os.Exit(), log.Fatal() or diary.Fatal()
type MessageHandler func(msg IMessage) (interface{}, error)