package types

type Password string

func (p Password) String() string { return string(p) }

func (p Password) SecurityString() string { return "--------" }
