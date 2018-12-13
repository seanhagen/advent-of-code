package day13

import "testing"

func TestCreateCart(t *testing.T) {
	tests := []struct {
		in  string
		inx int
		iny int
		f   Facing
	}{
		{">", 2, 3, FacingEast},
		{"<", 3, 4, FacingWest},
		{"^", 4, 5, FacingNorth},
		{"v", 1, 2, FacingSouth},
	}

	for _, tt := range tests {
		c := CreateCart(tt.in, tt.inx, tt.iny)

		if c.x != tt.inx {
			t.Errorf("wrong x -- expected %v, got %v", tt.inx, c.x)
		}

		if c.y != tt.iny {
			t.Errorf("wrong y -- expected %v, got %v", tt.iny, c.y)
		}

		if c.facing != tt.f {
			t.Errorf("wrong facing -- expected %#v, got %#v", tt.f, c.facing)
		}
	}
}

func TestCartMove(t *testing.T) {
	tests := []struct {
		in     string
		startx int
		starty int
		nextx  int
		nexty  int
	}{
		{">", 1, 1, 2, 1},
		{"^", 1, 1, 1, 2},
		{"v", 1, 1, 1, 0},
		{"<", 1, 1, 0, 1},
	}

	for _, tt := range tests {
		c := CreateCart(tt.in, tt.startx, tt.starty)
		c.Move()

		if c.x != tt.nextx {
			t.Errorf("wrong x -- expected %v, got %v", tt.nextx, c.x)
		}

		if c.y != tt.nexty {
			t.Errorf("wrong y -- expected %v, got %v", tt.nexty, c.y)
		}
	}
}

func TestCartNext(t *testing.T) {
	tests := []struct {
		in     []string
		startx int
		starty int
		endx   int
		endy   int
	}{
		// >-- => ->- => -->
		{[]string{">", "-", "-"}, 0, 0, 2, 0},
		// // |      |      ^
		// // |  =>  ^  =>  |
		// // ^      |      |
		{[]string{"^", "|", "|"}, 0, 0, 0, 2},
		// //  >-\   =>  ->\   =>  --v  =>  --\
		// //    |         |         |        v
		{[]string{">", "-", "\\", "|"}, 10, 10, 12, 9},
		// // >-\  => ->\  => --v
		{[]string{">", "-", "\\"}, 0, 0, 2, 0},

		// />-\    /->\    /--v    /--\    /--\    /--\    /--\    /--\    /--\    >--\    />-\
		// |  | => |  | => |  | => |  v => |  | => |  | => |  | => |  | => ^  | => |  | => |  |
		// \--/    \--/    \--/    \--/    \--<    \-</    \<-/    ^--/    \--/    \--/    \--/
		{[]string{">", "-", "\\", "|", "/", "-", "-", "\\", "|", "/", "-"}, 1, 1, 1, 1},

		// >-\  => ->\  => --v  => --\  => --\
		//   +-      +-      +-      >-      +-
		{[]string{">", "-", "\\", "+", "-"}, 0, 1, 3, 0},

		// 10,10     11,10     12,10    12,9   13,9     14,9        14,8
		// >-\    => ->\    => --v   => --\    => --\   => --\   => --\
		//   +++       +++       +++      >+++      +>+      ++v      +++
		//     |         |         |         |        |        |        v
		{[]string{">", "-", "\\", "+", "+", "+", "|"}, 10, 10, 14, 8},
	}

	for _, tt := range tests {
		i := tt.in[0]
		moves := tt.in[1:]

		c := CreateCart(i, tt.startx, tt.starty)

		for _, x := range moves {
			c.Process(x)
		}

		if tt.endx != c.x {
			t.Errorf("wrong endpoint x -- expected %v, got %v", tt.endx, c.x)
		}

		if tt.endy != c.y {
			t.Errorf("wrong endpoint y -- expected %v, got %v", tt.endy, c.y)
		}
	}
}
