package dice

import "testing"

func TestNewDice(t *testing.T) {
	t.Run("size=1", func(t *testing.T) {
		d1, err := NewDice(1)
		if err == nil || d1 != nil {
			t.Error("invalid dice size should return error, received nil")
		}
	})

	t.Run("size>1", func(t *testing.T) {
		d5, err := NewDice(5)

		if err != nil {
			t.Errorf("error must be nil, received: %e", err)
		}

		if d5.size != 5 {
			t.Errorf("dice size should be 5, received %d", d5.size)
		}
	})

}

func TestDiceRoll(t *testing.T) {
	t.Run("result!=0", func(t *testing.T) {
		d2, _ := NewDice(2)
		r := d2.Roll()
		if r == 0 {
			t.Errorf("result should be bigger than zero, received %d", r)
		}

		r = d2.Roll()
		if r == 0 {
			t.Errorf("result should be bigger than zero, received %d", r)
		}

		r = d2.Roll()
		if r == 0 {
			t.Errorf("result should be bigger than zero, received %d", r)
		}
	})

	t.Run("result!=0&&<=size", func(t *testing.T) {
		d2, _ := NewDice(2)
		r := d2.Roll()
		if r == 0 || r > 2 {
			t.Errorf("result should be between zero and %d, received %d", d2.size, r)
		}

		r = d2.Roll()
		if r == 0 || r > 2 {
			t.Errorf("result should be between zero and %d, received %d", d2.size, r)
		}

		r = d2.Roll()
		if r == 0 || r > 2 {
			t.Errorf("result should be between zero and %d, received %d", d2.size, r)
		}
	})
}
