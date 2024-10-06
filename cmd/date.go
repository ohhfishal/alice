package cmd

type date string

func (d date) String() string {
  return (string)(d)
}

func (d date) Set(string) error {
  // TODO: Validate its something like "tomorrow@noon" or an actual date
  return nil
}

func (d date) Type() string {
  return "date"
}
