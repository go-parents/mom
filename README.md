# Message-Oriented Middle-ware
A universal Message-Oriented Middle-ware (MOM) connection interface.

NewNatsConnection() IConn

type IConn interface{
    Publish()
    Request()
    Subscribe()

    StreamIn()
    StreamOut()
}
