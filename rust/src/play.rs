use world::World;
use time;

pub struct Play {}

impl Play {

  const World_Width: u32 = 150;
  const World_Height: u32 = 40;

  pub fn run() {
    let mut world = World::new(
      Play::World_Width,
      Play::World_Height
    );

    println!("{}", world.render());

    let mut total_tick: f64 = 0.0;
    let mut total_render: f64 = 0.0;

    while true {
      let tick_start = time::precise_time_s();
      world._tick();
      let tick_finish = time::precise_time_s();
      let tick_time = (tick_finish - tick_start) / 1000.0;
      total_tick += tick_time;
      let avg_tick = total_tick / world.tick as f64;

      let render_start = time::precise_time_s();
      let rendered = world.render();
      let render_finish = time::precise_time_s();
      let render_time = (render_finish - render_start) / 1000.0;
      total_render += render_time;
      let avg_render = total_render / world.tick as f64;

      let mut output = format!("#{}", world.tick);
      output += &format!(" - World tick took {} ({})", Play::_f(tick_time), Play::_f(avg_tick));
      output += &format!(" - Rendering took {} ({})", Play::_f(render_time), Play::_f(avg_render));
      output += &format!("\n{}", rendered.to_string());
      println!("{}", "\033[H\033[2J");
      println!("{}", output);
    }
  }

  pub fn _f(value: f64) -> String {
    return format!("{:.5}", value);
  }

}
