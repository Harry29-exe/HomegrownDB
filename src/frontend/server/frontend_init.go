package server

//type DBServer struct {
//	listener net.Listener
//	stop     chan bool
//}

//func StartNetworkInterface(address string, port string) (DBServer, error) {
//	portListener, err := net.Listen("tcp", address+":"+port)
//	if err != nil {
//		return DBServer{}, err
//	}
//
//}

//func (s *DBServer) listen() {
//	for {
//		select {
//		case <-s.stop:
//			fmt.Println("Closing db server")
//			return
//		default:
//			conn, err := s.listener.Accept()
//			if err != nil {
//				continue
//			}
//		}
//	}
//}

//func (s *DBServer) Close() error {
//	s.stop <- true
//	err := s.listener.Close()
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
