import 'dart:async';
import 'dart:convert';
import 'dart:html';

import '../pathlookup.dart';

Future<HttpRequest> createComment(
    String itemKey, String text, String commentType) async {
  var url = await buildPath("Comment.API", "message", new List<String>());
  var data = jsonEncode({
    "ItemKey": itemKey,
    "Text": text,
    "CommentType": commentType,
    "Children": []
  });

  final compltr = new Completer<HttpRequest>();
  final request = HttpRequest();
  request.open("POST", url);
  request.setRequestHeader('Content-Type', 'application/json; charset=utf-8');
  request.setRequestHeader(
      "Authorization", "Bearer " + window.localStorage['avosession']);
  request.onLoadEnd
      .listen((e) => compltr.complete(request), onError: compltr.completeError);
  request.onError.listen(compltr.completeError);
  request.onProgress.listen(onProgress);
  request.send(data);

  return compltr.future;
}

void onProgress(ProgressEvent e) {
  if (e.lengthComputable) {
    print('Progress... ${e.total}/${e.loaded}');
  }
}
