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
