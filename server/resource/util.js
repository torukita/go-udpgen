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

Client.prototype.postConfig = function(json) {
    return this.client.POST(this.baseUrl + "config", json)
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
