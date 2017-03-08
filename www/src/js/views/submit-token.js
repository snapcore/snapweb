var Backbone = require('backbone');
var Marionette = require('backbone.marionette');
var template = require('../templates/submit-token.hbs');

module.exports = Backbone.Marionette.ItemView.extend({

  className: 'b-layout__container',

  template: function(model) {
    return template(model);
  },

  serializeData: function() {
    return {
      'appName': (window.location.port === '5201' ? 'ubuntu-personal-store' : 'snapweb'),
    };
  },

  ui: {
    statusmessage: '.statusmessage',
    btncreate: '#submit',
  },

  events: {
    'click #submit': 'handleSubmit',
  },

  modelEvents: {
    'invalid': function(model, error) {
      this.setErrorStatus(error);
    },
  },

  setErrorStatus: function(msg) {
    this.ui.statusmessage.text(msg);
    this.ui.statusmessage.removeClass('has-warning');
    this.ui.statusmessage.addClass('has-error');
    this.ui.statusmessage.show();
  },

  handleSubmit: function(event) {
    event.preventDefault();
    event.stopPropagation();
    this.model.setCookie(this.$('#token').val());
    this.model.save({}, {
      success: function() {
        // redirect to home for now
        window.location = (window.location.port === '5201' ? '/store': '/');
      },
      error: function(model, response) {
        model.trigger('invalid', model, 'Invalid');
      }
    });
  },

});
