import 'dart:html';

void main() {
  print('Master Theme Loaded.');
  enableTabs();
  enableBurger();
}

void enableTabs() {
  querySelectorAll("#nav li")
      .forEach((navEl) => {navEl.onClick.listen(tabClick)});
}

void tabClick(Event e) {
  Element elem = e.currentTarget;
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