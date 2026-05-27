package core

type Opcode uint8

const (
	ModuleOpcodeEOF Opcode = iota
	ModuleOpcodeSInt
	ModuleOpcodeUInt
	ModuleOpcodeFloat
	ModuleOpcodeDouble
	ModuleOpcodeString
)

type ModuleTypeHandler interface {
	ReadByte() (byte, error)
	ReadFull(buf []byte) error
	ReadOpcode() (Opcode, error)
	ReadUInt() (uint64, error)
	ReadSInt() (int64, error)
	ReadFloat32() (float32, error)
	ReadDouble() (float64, error)
	ReadString() ([]byte, error)
	ReadLength() (uint64, bool, error)
}

type moduleTypeHandlerImpl struct {
	dec *Decoder
}

func (m moduleTypeHandlerImpl) ReadByte() (byte, error) { _ = "STUB: not implemented"; return 0, nil }

func (m moduleTypeHandlerImpl) ReadFull(buf []byte) error { _ = "STUB: not implemented"; return nil }

func (m moduleTypeHandlerImpl) ReadOpcode() (Opcode, error) {
	_ = "STUB: not implemented"
	return *new(Opcode), nil
}

func (m moduleTypeHandlerImpl) ReadUInt() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (m moduleTypeHandlerImpl) ReadSInt() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (m moduleTypeHandlerImpl) ReadFloat32() (float32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m moduleTypeHandlerImpl) ReadDouble() (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m moduleTypeHandlerImpl) ReadString() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m moduleTypeHandlerImpl) ReadLength() (uint64, bool, error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

type ModuleTypeHandleFunc func(handler ModuleTypeHandler, encVersion int) (interface{}, error)

func (dec *Decoder) readModuleType() (string, interface{}, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func (dec *Decoder) handleModuleType(moduleId uint64) (string, interface{}, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func moduleTypeNameByID(moduleId uint64) string { _ = "STUB: not implemented"; return "" }

func moduleTypeEncVersionByID(moduleId uint64) uint64 { _ = "STUB: not implemented"; return 0 }

// skipModuleAuxData skips module aux data
func skipModuleAuxData(h ModuleTypeHandler, _ int) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const ModuleTypeNameCharSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
