package messaging

type Channel struct {
	C       chan *Message
	enabled bool
}

func NewChannel() *Channel {
	c := new(Channel)
	c.C = make(chan *Message, ChannelSize)
	return c
}

func NewEnabledChannel() *Channel {
	c := NewChannel()
	c.enabled = true
	return c
}

func (c *Channel) IsEnabled() bool {
	return c.enabled
}

func (c *Channel) Enable() {
	c.enabled = true
}

func (c *Channel) Disable() {
	c.enabled = false
}

func (c *Channel) Close() {
	if c.C != nil {
		close(c.C)
		c.C = nil
	}
}

func (c *Channel) Send(m *Message) {
	if m != nil && c.enabled {
		c.C <- m
	}
}
