package message

type Event string

const Snapshot Event = "snapshot"
const Insert Event = "insert"
const Update Event = "update"
const Delete Event = "delete"
