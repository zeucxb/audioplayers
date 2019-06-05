import 'package:flutter/foundation.dart'
    show debugDefaultTargetPlatformOverride;

import 'package:flutter/material.dart';
import './example_app.dart';

typedef void OnError(Exception exception);

void main() {
  debugDefaultTargetPlatformOverride = TargetPlatform.fuchsia;
  runApp(new MaterialApp(home: new ExampleApp()));
}

