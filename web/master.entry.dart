import 'dart:html';

void main() {
  enableTabs();
  enableBurger();
}

void enableTabs() {
  document.body.onClick.matches("#nav li").listen(tabClick);
}

void tabClick(Event e) {
  Element elem = e.matchingTarget;
  toggleTab(elem.id, elem.dataset["target"]);
}

void toggleTab(selectedNav, targetId) {
  var navEls = document.querySelectorAll("#nav li");

  navEls.forEach((navEl) {
    if (navEl.id == selectedNav) {
      navEl.classes.add("is-active");
    } else {
      if (navEl.classes.contains("is-active")) {
        navEl.classes.remove("is-active");
      }
    }
  });

  var tabs = document.querySelectorAll(".tab-pane");

  tabs.forEach((tab) {
    if (tab.id == targetId) {
      tab.style.display = "block";
    } else {
      tab.style.display = "none";
    }
  });
}

void enableBurger() {
  var burger = querySelector('.burger');
  var menu = document.querySelector('#${burger.dataset["target"]}');
  burger.onClick.listen((e) => burgerClick(burger, menu));
}

void burgerClick(Element burger, Element menu) {
  burger.classes.toggle('is-active');
  menu.classes.toggle('is-active');
}

String getParameterByName(String name) {
  final url = window.location.href;
  var uri = Uri.parse(url);
  return uri.queryParameters[name];
}
