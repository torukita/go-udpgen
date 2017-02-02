package resource

const(
        TemplateText=`
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta name="description" content="">
  <meta name="author" content="">
  <title>Simple UDP generator</title>
  <!-- Bootstrap core CSS -->
  <link href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
  <!-- Custom styles for this template -->
</head>
<body>
  <h1>Simple UDP Generator</h1>
  <div class="container">
    <div class="row"><div class="col-md-12"><div id=status></div></div>
    <div class="row">
      <div class="col-md-3">
        <div class="btn-group" role="group" aria-label="...">
          <div class="btn-group">
            <button type="button" class="btn btn-primary dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">devices <span class="caret"></span></button>
            <ul class="dropdown-menu device">
              {{range .}}
              <li><a href="#{{.Index}}" data-value="{{.MacAddr}}">{{.Name}}</a></li>
              {{end}}
            </ul>
          </div>
        </div>
      </div>
      <div class="col-md-9"></div>
    </div>
    <div class="row">
      <div class="col-md-6">
        <div class="input-group">
          <span class="input-group-addon" id="sizing-addon1">Src Address</span>
          <input type="text" class="form-control" id="eth-src" placeholder="src mac" aria-describedby="sizing-addon1">
        </div>
      </div>
      <div class="col-lg-6">
        <div class="input-group">
          <div class="input-group-btn">
            <button type="button" class="btn btn-default dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">Dst Address <span class="caret"></span></button>
            <ul class="dropdown-menu eth-menu">
              <li><a href="#" data-value="hogehoge"></a></li>
            </ul>
          </div><!-- /btn-group -->
          <input type="text" class="form-control" id="eth-dst" aria-label="...">
        </div><!-- /input-group -->
      </div><!-- /.col-lg-6 -->
    </div>
    <div class="row">
      <div class="col-md-6">
        <div class="input-group">
          <span class="input-group-addon" id="sizing-addon1">Src IP</span>
          <input type="text" class="form-control" id="ip-src" placeholder="10.0.0.1" aria-describedby="sizing-addon1">
        </div>
      </div>
      <div class="col-md-6">
        <div class="input-group">
          <span class="input-group-addon" id="sizing-addon1">Dst IP</span>
          <input type="text" class="form-control" id="ip-dst" placeholder="10.0.0.2" aria-describedby="sizing-addon1">
        </div>
      </div><!-- /.col-lg-6 -->
    </div>
    <div class="row">
      <div class="col-md-6">
        <div class="input-group">
          <span class="input-group-addon" id="sizing-addon1">Src PORT</span>
          <input type="text" class="form-control" id="udp-src" placeholder="5000" aria-describedby="sizing-addon1">
        </div>
      </div>
      <div class="col-md-6">
        <div class="input-group">
          <span class="input-group-addon" id="sizing-addon1">Dst PORT</span>
          <input type="text" class="form-control" id="udp-dst" placeholder="5000" aria-describedby="sizing-addon1">
        </div>
      </div><!-- /.col-lg-6 -->
    </div>
  </div>

  <script src="https://code.jquery.com/jquery-3.1.1.min.js"></script>
  <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>
  <script src="/js/utils.js"></script>
  <script>

// initial function after document load
$(function() {
    $(".eth-menu li").remove()
    for (var i=0; i< ethdst.length; i++) {
        $(".eth-menu").append('<li><a href="#" data-value="' + ethdst[i] + '" onclick=test(this);>' + ethdst[i] + '</a></li>');
    }
});

var ethdst = ["00:00:00:00:00:01", "00:00:00:00:00:02", "00:00:00:00:00:03"];

var notice = new Notifier({selector: "#status"});
$(".device li a").click(function() {
    var device = $(this).attr('data-value');
    notice.success("Done:" + device);
    $("#eth-src").attr("placeholder", device);
});

function test(data) {
    var value = data.getAttribute("data-value");
    notice.success("Done:" + value);
    $("#eth-dst").attr("placeholder", value);
};

  </script>
</body>
</html>
`
      JSText=`
Notifier = function(options) {
    var defaultOpts = {
        selector: "#demo-status"
    };
    var config = $.extend({}, defaultOpts, options);
    this.selector = config.selector;
};

Notifier.prototype._generate = function(status, message) {
    var block = '<div class="alert alert-' + status + ' alert-dismissible" role="alert">';
    block = block + '<button type="button" class="close" data-dismiss="alert" aria-label="Close">';
    block = block + '<span aria-hidden="true">&times;</span></button>' + message + '</div>';
    return block;
};

Notifier.prototype.success = function(message) {
    $(this.selector).html(this._generate("success", message));
};
Notifier.prototype.info = function(message) {
    $(this.selector).html(this._generate("info", message));
};
Notifier.prototype.warning = function(message) {
    $(this.selector).html(this._generate("warning", message));
};
Notifier.prototype.danger = function(message) {
    $(this.selector).html(this._generate("danger", message));
};
`
)
