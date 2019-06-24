import 'dart:convert';
import 'dart:html';

import 'package:Theme.API/formstate.dart';
import 'package:Theme.API/services/commentapi.dart';

class Comments extends FormState {
  String _objKey;
  String _objType;
  TextInputElement _text;

  Comments(String idElem, String objKey, String objType)
      : super(idElem, "#btnComment") {
    _objKey = objKey;
    _objType = objType;
    _text = querySelector("#txtText");

    //data-itemKey="$key" data-itemType="Child"
    querySelector("#btnComment").onClick.listen(onCommentClick);
  }

  String get text {
    return _text.value;
  }

  void onCommentClick(MouseEvent e) async {
    if (isFormValid()) {
      disableSubmit(true);

      var req = await createComment(_objKey, text, _objType);

      final resp = jsonDecode(req.response);

      if (req.status == 200) {
        window.alert(resp['Data']);
      }
    }
  }
}
