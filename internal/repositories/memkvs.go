package repositories

import ()

type memkvs struct {}

func LoadMemKVS() *memkvs {
	var memKVS memkvs
	return &memKVS
}

