package products

//----- mocking store
type StubStore struct {
	Data []Product
	errOnWrite error
	errOnRead error

}

func (m *StubStore) Read(data interface{}) error {
	if m.errOnRead != nil {
		return m.errOnRead
	}
	castedData := data.(*[]Product)
	*castedData = m.Data
	return nil
}

func (m *StubStore) Write(data interface{}) error {
	if m.errOnWrite != nil {
		return m.errOnWrite
	}
	castedData := data.([]Product)
	m.Data = append(m.Data, castedData[len(castedData)-1])	
	// castedData := data.(*Product)
	// m.data = append(m.data, *castedData)	 No lo toma
	return nil
}
