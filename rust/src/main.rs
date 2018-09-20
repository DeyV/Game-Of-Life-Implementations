extern crate rand;
extern crate time;

mod world;
mod play;

use play::Play;

fn main() {
  Play::run();
}
