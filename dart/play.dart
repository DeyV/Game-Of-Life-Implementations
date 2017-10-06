// In Dart, anything preceeded by an underscore (_) is a private function/variable

import 'world.dart';
import 'dart:io';

class Play {

  static const World_Width  = 150;
  static const World_Height = 40;

  static void run() {
    final world = new World(
      World_Width,
      World_Height,
    );

    print(world.render());

    var total_tick = 0.0;
    var total_render = 0.0;

    while(true) {
      final tick_stopwatch = new Stopwatch()..start();
      final tick_start = tick_stopwatch.elapsedMilliseconds;
      world.tick_();
      final tick_finish = tick_stopwatch.elapsedMilliseconds;
      final tick_time = (tick_finish - tick_start) / 1000.0;
      total_tick += tick_time;
      final avg_tick = (total_tick / world.tick);

      final render_stopwatch = new Stopwatch()..start();
      final render_start = render_stopwatch.elapsedMilliseconds;
      final rendered = world.render();
      final render_finish = render_stopwatch.elapsedMilliseconds;
      final render_time = (render_finish - render_start) / 1000.0;
      total_render += render_time;
      final avg_render = (total_render / world.tick);

      var output = "#${world.tick}";
      output += " - World tick took ${_f(tick_time)} (${_f(avg_tick)})";
      output += " - Rendering took ${_f(render_time)} (${_f(avg_render)})";
      output += "\n${rendered}";
      stdout.write("\u001b[H\u001b[2J");
      stdout.write(output);
    }
  }

  static String _f(int value) {
    return value.toStringAsFixed(5);
  }

}

main() {
  Play.run();
}