package gonbt

// Tag Name
func TagNamePtr(x TagName) *TagName { return &x }

// Payload
func BytePayloadPtr(x BytePayload) *BytePayload       { return &x }
func ShortPayloadPtr(x ShortPayload) *ShortPayload    { return &x }
func IntPayloadPtr(x IntPayload) *IntPayload          { return &x }
func LongPayloadPtr(x LongPayload) *LongPayload       { return &x }
func FloatPayloadPtr(x FloatPayload) *FloatPayload    { return &x }
func DoublePayloadPtr(x DoublePayload) *DoublePayload { return &x }
func StringPayloadPtr(x StringPayload) *StringPayload { return &x }
