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
    <div class="row">
      <div class="col-md-12"><div id="status"></div></div></div>
    <div class="row">
      <div class="col-md-3">
        <div class="btn-group" role="group" aria-label="...">
          <div class="btn-group">
            <button type="button" class="btn btn-primary dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">devices <span class="caret"></span></button>
            <ul class="dropdown-menu" id="device-list"></ul>
          </div>
        </div>
      </div>
      <div class="col-md-9"></div>
    </div>
    <br/>
    <form id="udpgen-config">
    <div class="row">
      <div class="col-md-6">
        <div class="input-group">
          <span class="input-group-addon" id="sizing-addon1">Src PORT </span>
          <input type="text" class="form-control" id="src-udp" name="SrcPort" value="5000" aria-describedby="sizing-addon1">
        </div>
      </div><!-- src port -->
      <div class="col-md-6">
        <div class="input-group">
          <span class="input-group-addon" id="sizing-addon1">Dst PORT </span>
          <input type="text" class="form-control" id="dst-udp" name="DstPort" value="5000" aria-describedby="sizing-addon1">                   </div>
      </div><!-- dst port -->
    </div>
    <div class="row">
      <div class="col-md-6">
        <div class="input-group">
          <span class="input-group-addon" id="sizing-addon1">Src Addr </span>
          <input type="text" class="form-control" id="src-eth" name="SrcEth" placeholder="Source MacAddress" aria-describedby="sizing-addon1">
        </div>
      </div><!-- src eth -->
      <div class="col-md-6">
        <div class="input-group">
          <div class="input-group-btn">
            <button type="button" class="btn btn-default dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">Dst Addr <span class="caret"></span></button>
            <ul class="dropdown-menu" id="dst-eth-list"></ul>
          </div>
          <input type="text" class="form-control" id="dst-eth" name="DstEth" placeholder="Destination MacAddress" aria-label="...">
        </div>
      </div>
    </div><!-- dst eth -->
    <div class="row">
      <div class="col-md-6">
        <div class="input-group">
          <span class="input-group-addon" id="sizing-addon1">Src IP </span>
          <input type="text" class="form-control" id="src-ip" name="SrcIP" placeholder="Source IP" aria-describedby="sizing-addon1">
        </div>
      </div><!-- src ip -->
      <div class="col-md-6">
        <div class="input-group">
          <div class="input-group-btn">
            <button type="button" class="btn btn-default dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">Dst IP <span class="caret"></span></button>
            <ul class="dropdown-menu" id="dst-ip-list">
            </ul>
          </div>
          <input type="text" class="form-control" id="dst-ip" name="DstIP" placeholder="Destination IP" aria-label="...">
        </div>
      </div><!-- dst ip -->
    </div>
    <br/>
    <div class="row">
      <div class="col-md-3">
        <button type="button" class="btn btn-primary" id="form-post")>GO!</button>
      </div>
      <div class="col-md-9"></div>
    </div>
    </form>
  </div>

  <script src="https://code.jquery.com/jquery-3.1.1.min.js"></script>
  <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>
  <script src="/js/utils.js"></script>
  <script>

// initial function after document load
$(function() {
    DeviceList();
});

var notice = new Notifier({selector: "#status"});
var api = new Client();

function DeviceList() {
    api.getDevices().done(function (devices) {
        var list = []
        for (var i=0; i < devices.length; i++) {
            list.push('<li><a href="#" class="device" deviceIndex="' + devices[i].Index + '" macAddr="' + devices[i].MacAddr
                      + '">' + devices[i].Name + '</a></li>');
        }
        $('#device-list').html(list);
    });
}

$('#device-list').on("click", "a.device", function(evt) {
    evt.preventDefault();
    var index = $(this).attr('deviceIndex');
    var macAddr = $(this).attr('macAddr');
    var ip = "10.0.0.1"; // default IP
    $('#src-eth').val(macAddr);
    api.getIPByIndex(index).done(function(data) {
        if (data.length > 0) {
            ip = data[0].IP; // pick up first ip
        }
        $('#src-ip').val(ip);
    });
    UpdateDstEtherAndIPMenu(index);
});

function UpdateDstEtherAndIPMenu(index) {
    api.getArpTable(index).done(function(tables) {
        var ipList = [];
        var macList = [];
        for (var i=0; i< tables.length; i++) {
            var mac = tables[i].MacAddr;
            var ip = tables[i].IP;
            macList.push('<li><a href="#" class="dsteth" macAddr="' + mac + '">' + mac + ' ( ' + ip + ')</a></li>');
            ipList.push('<li><a href="#" class="dstip" ip="' + ip + '">' + ip + ' ( ' + mac + ')</a></li>');
        }
        $('#dst-eth-list').html(macList);
        $('#dst-ip-list').html(ipList);
    });
}

$('#dst-eth-list').on("click", "a.dsteth", function(evt) {
    evt.preventDefault();
    var macAddr = $(this).attr("macAddr");
    $('#dst-eth').val(macAddr);
});

$('#dst-ip-list').on("click", "a.dstip", function(evt) {
    evt.preventDefault();
    var ip = $(this).attr("ip");
    $('#dst-ip').val(ip);
});

$('#form-post').on('click', function() {
    var data = JSON.stringify( $('#udpgen-config').serializeArray());
    console.log(data);
});

</script>
</body>
</html>
`
      JSText=`
Client = function(options) {
    this.baseUrl = "http://" + window.location.hostname + ":9000/api/";
    this.client = new Rest();
};

Client.prototype.getDevices = function() {
    return this.client.GET(this.baseUrl + "devices")
};

Client.prototype.getArpTable = function(index) {
    return this.client.GET(this.baseUrl + "device/" + index + "/arp")
};

Client.prototype.getIPByIndex = function(index) {
    return this.client.GET(this.baseUrl + "device/" + index + "/ipv4")
};

Client.prototype.postFlow = function(json) {
    return this.client.POST(this.baseUrl + "/flow", json)
};

// Never use this class from user program
// This is under construction.
// TODO: better to implement without jquery dependency in the future
Rest = function(options) {
    if (options != null) {
    }
    // this overrides jquery ajax as default
    $.ajaxSetup({
        contentType: 'application/json',
        dataType: 'json',
        jsonp: false
    });
};

Rest.prototype.GET = function(url, json) {
    return $.ajax({
        type: 'get',
        url: url,
        data: json
    })
};

Rest.prototype.POST = function(url, json) {
    return $.ajax({
        type: 'post',
        url: url,
        data: json
    })
};

Rest.prototype.DELETE = function(url, json) {
    return $.ajax({
        type:'delete',
        url: url,
        data: json
    })
};

Rest.prototype.PUT = function(url, json) {
    return $.ajax({
        type:'put',
        url: url,
        data: json
    })
};

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
