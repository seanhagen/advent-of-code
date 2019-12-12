package day12

// CalculateVelocity calculates the velocity for a pair of planets and modifies their
// stored velocities
func CalculateVelocity(p1, p2 *Planet) {
	p1vd, p2vd := Vec{1, 1, 1}, Vec{1, 1, 1}

	if p1.pos.x != p2.pos.x {
		if p1.pos.x < p2.pos.x {
			p2vd.x *= -1
		} else {
			p1vd.x *= -1
		}
	} else {
		p1vd.x = 0
		p2vd.x = 0
	}

	if p1.pos.y != p2.pos.y {
		if p1.pos.y < p2.pos.y {
			p2vd.y *= -1
		} else {
			p1vd.y *= -1
		}
	} else {
		p1vd.y = 0
		p2vd.y = 0
	}

	if p1.pos.z != p2.pos.z {
		if p1.pos.z < p2.pos.z {
			p2vd.z *= -1
		} else {
			p1vd.z *= -1
		}
	} else {
		p1vd.z = 0
		p2vd.z = 0
	}

	p1.vel.Add(p1vd)
	p2.vel.Add(p2vd)
}
