package gonbt

type TagType byte

const (
	TagEnd TagType = iota
	TagByte
	TagShort
	TagInt
	TagLong
	TagFloat
	TagDouble
	TagByteArray
	TagString
	TagList
	TagCompound
	TagIntArray
	TagLongArray
)

// Payload
func (p *BytePayload) TypeId() TagType      { return TagByte }
func (p *ShortPayload) TypeId() TagType     { return TagShort }
func (p *IntPayload) TypeId() TagType       { return TagInt }
func (p *LongPayload) TypeId() TagType      { return TagLong }
func (p *FloatPayload) TypeId() TagType     { return TagFloat }
func (p *DoublePayload) TypeId() TagType    { return TagDouble }
func (p *ByteArrayPayload) TypeId() TagType { return TagByteArray }
func (p *StringPayload) TypeId() TagType    { return TagString }
func (p *ListPayload) TypeId() TagType      { return TagList }
func (p *CompoundPayload) TypeId() TagType  { return TagCompound }
func (p *IntArrayPayload) TypeId() TagType  { return TagIntArray }
func (p *LongArrayPayload) TypeId() TagType { return TagLongArray }

// Tag
func (t *EndTag) TypeId() TagType       { return TagEnd }
func (t *ByteTag) TypeId() TagType      { return TagByte }
func (t *ShortTag) TypeId() TagType     { return TagShort }
func (t *IntTag) TypeId() TagType       { return TagInt }
func (t *LongTag) TypeId() TagType      { return TagLong }
func (t *FloatTag) TypeId() TagType     { return TagFloat }
func (t *DoubleTag) TypeId() TagType    { return TagDouble }
func (t *ByteArrayTag) TypeId() TagType { return TagByteArray }
func (t *StringTag) TypeId() TagType    { return TagString }
func (t *ListTag) TypeId() TagType      { return TagList }
func (t *CompoundTag) TypeId() TagType  { return TagCompound }
func (t *IntArrayTag) TypeId() TagType  { return TagIntArray }
func (t *LongArrayTag) TypeId() TagType { return TagLongArray }
