import 'dart:html';

void main() {
  enableTabs();
  enableBurger();
}

void enableTabs() {
  document.body.onClick.matches("#nav li").listen(tabClick);
  document.body.onClick.matches(".card-header").listen(collapseClick);
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

void collapseClick(Event e) {
  Element elem = e.matchingTarget;
  Element content = elem.nextElementSibling;

  if (content != null) {
    Element icon = elem.children
        .where((element) => element.classes.contains("card-header-icon"))
        .first;
    Element img = icon.children.first.children.first;
    if (content.hidden) {
      img.classes.remove('fa-angle-down');
      img.classes.add('fa-angle-up');
    } else {
      img.classes.remove('fa-angle-up');
      img.classes.add('fa-angle-down');
    }

    content.hidden = !content.hidden;
  }
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
