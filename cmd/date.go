package cmd

type _Date string

func (d _Date) String() string {
  return (string)(d)
}

func (d _Date) Set(string) error {
  // TODO: Vali_Date its something like "tomorrow@noon" or an actual _Date
  return nil
}

func (d _Date) Type() string {
  return "_Date"
}
