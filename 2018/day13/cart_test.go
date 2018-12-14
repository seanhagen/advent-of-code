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
		{"^", 1, 1, 1, 0},
		{"v", 1, 1, 1, 2},
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
		{[]string{"^", "|", "|"}, 0, 2, 0, 0},
		// //  >-\   =>  ->\   =>  --v  =>  --\
		// //    |         |         |        v
		{[]string{">", "-", "\\", "|"}, 10, 10, 12, 11},
		// // >-\  => ->\  => --v
		{[]string{">", "-", "\\"}, 0, 0, 2, 0},

		// />-\    /->\    /--v    /--\    /--\    /--\    /--\    /--\    /--\    >--\    />-\
		// |  | => |  | => |  | => |  v => |  | => |  | => |  | => |  | => ^  | => |  | => |  |
		// \--/    \--/    \--/    \--/    \--<    \-</    \<-/    ^--/    \--/    \--/    \--/
		{[]string{">", "-", "\\", "|", "/", "-", "-", "\\", "|", "/", "-"}, 1, 1, 1, 1},

		// >-\  => ->\  => --v  => --\  => --\
		//   +-      +-      +-      >-      +>
		{[]string{">", "-", "\\", "+", "-"}, 0, 1, 3, 2},

		// 10,10     11,10     12,10    12,9   13,9     14,9        14,8
		// >-\    => ->\    => --v   => --\    => --\   => --\   => --\
		//   +++       +++       +++      >+++      +>+      ++v      +++
		//     |         |         |         |        |        |        v
		{[]string{">", "-", "\\", "+", "+", "+", "|"}, 10, 10, 14, 12},

		// 10,10  11,10  12,10  12,9
		//   | =>   | =>   | =>   ^
		// >-/    ->/    --^    --/
		{[]string{">", "-", "/", "|"}, 10, 10, 12, 9},

		// 10,10    9,10    8,10     8,11
		// /-<  =>  /<-  => v--   => /--
		// |        |       |        v
		{[]string{"<", "-", "/", "|"}, 10, 10, 8, 11},
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

func TestCartUnder(t *testing.T) {
	tests := []struct {
		in string
		u  string
	}{
		{"^", "|"},
		{">", "-"},
		{"v", "|"},
		{"<", "-"},
	}

	for _, tt := range tests {
		c := CreateCart(tt.in, 0, 0)
		g := c.Under()

		if tt.u != g {
			t.Errorf("wrong output, expected '%v' got '%v'", tt.u, g)
		}
	}
}

func TestCartNextPos(t *testing.T) {
	tests := []struct {
		in    string
		x     int
		y     int
		nextx int
		nexty int
	}{
		{">", 0, 0, 1, 0},
		{"^", 1, 1, 1, 0},
		{"v", 1, 1, 1, 2},
		{"<", 1, 1, 0, 1},
	}

	for _, tt := range tests {
		c := CreateCart(tt.in, tt.x, tt.y)
		nx, ny := c.NextPos()
		if tt.nextx != nx || tt.nexty != ny {
			t.Errorf("wrong next, expected <%v, %v>, got <%v, %v>", tt.nextx, tt.nexty, nx, ny)
		}
	}
}
