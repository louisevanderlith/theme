import "dart:html";

void main() {
  registerEvents();
}

void registerEvents() {
  var offcanvas = querySelector("[data-toggle=offcanvas]");

  offcanvas.onClick.listen(click_OffCanvas);
}

void click_OffCanvas(Event e) {
  e.matchingTarget.classes.toggle("visible-xs text-center");
  e.matchingTarget.children.first.classes
      .toggle("glyphicon-chevron-right glyphicon-chevron-left");

  querySelector(".row-off-canvas").classes.toggle("active");

  var lgMenu = querySelector("#lg-menu");

  lgMenu.classes.toggle("hidden-xs");
  lgMenu.classes.toggle("visible-xs");

  var xsMenu = querySelector("#xs-menu");
  xsMenu.classes.toggle("visible-xs");
  xsMenu.classes.toggle("hidden-xs");

  var btnShow = querySelector("#btnShow");

  btnShow.hidden = !btnShow.hidden;
}
