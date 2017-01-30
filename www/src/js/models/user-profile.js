var Backbone = require('backbone');
var Marionette = require('backbone.marionette');
var CONFIG = require('../config.js');

module.exports = Backbone.Model.extend({
  url: CONFIG.USER_PROFILE,

  defaults: {
    name: '',
    email: '',
    avatar: '<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" height="48px" width="48px" version="1.1" viewBox="0 0 15 15"><g id="Page-1" fill-rule="evenodd" fill="none"><g id="profile" fill="#808080"><path id="path4116" d="m7.5 1.7764e-15c-4.1379 0-7.5 3.365-7.5 7.5059 1.1102e-13 4.1411 3.3621 7.5061 7.5 7.5061 4.138 0 7.5-3.365 7.5-7.5061 0-4.1409-3.362-7.5059-7.5-7.5059v1.7764e-15zm0 0.71485c3.752 0 6.785 3.0362 6.785 6.791 0 1.6665-0.598 3.1902-1.59 4.3712-0.216-0.498-0.51-0.947-0.882-1.339-0.35-0.369-0.77-0.6774-1.254-0.928-0.204 0.1799-0.426 0.3395-0.6598 0.479 0.5698 0.244 1.0318 0.557 1.3948 0.941 0.386 0.406 0.67 0.876 0.859 1.419-1.214 1.145-2.85 1.848-4.653 1.848-1.8098 0-3.4509-0.708-4.6666-1.861 0.2005-0.538 0.4832-1.005 0.8535-1.407 0.3943-0.416 0.9027-0.75 1.5405-1.002 0.6262-0.2467 1.3806-0.3763 2.267-0.3769 0.0019 0 0.0037 0.0004 0.0056 0.0004 0.5154 0 1.0023-0.0907 1.4502-0.2754l0.0028-0.0014 0.0031-0.0018c0.4445-0.1929 0.8322-0.4676 1.1519-0.8178 0.328-0.3505 0.582-0.7688 0.757-1.245 0.178-0.4814 0.264-1.0137 0.264-1.5903v-0.0003c0-0.5685-0.086-1.096-0.264-1.5767-0.175-0.4836-0.428-0.9066-0.757-1.2576-0.3197-0.3511-0.7092-0.6223-1.1554-0.8067-0.4483-0.194-0.9362-0.2904-1.4526-0.2904-0.5163 0-1.0043 0.0964-1.4526 0.2904-0.4463 0.1844-0.8394 0.4545-1.1677 0.8042l-0.0018 0.0021-0.0021 0.0021c-0.3199 0.3515-0.567 0.7741-0.7418 1.2573l-0.0004 0.0007c-0.1767 0.4802-0.2626 1.0069-0.2626 1.5746 0 0.5766 0.0859 1.1092 0.2633 1.5906 0.1751 0.4748 0.4225 0.8921 0.7415 1.2426l0.0021 0.0021 0.0018 0.0017c0.2193 0.2337 0.4679 0.433 0.7411 0.5976-0.2274 0.0581-0.4464 0.1269-0.655 0.2091-0.7195 0.2836-1.324 0.6748-1.7986 1.1758l-0.0017 0.002-0.0018 0.002c-0.3585 0.389-0.6468 0.833-0.867 1.325-0.9873-1.18-1.5821-2.6997-1.5821-4.3611 0.00003-3.7548 3.0337-6.791 6.7854-6.791v-0.00005zm0 1.7868c0.4299 0 0.8166 0.0783 1.1719 0.2325l0.0031 0.0014 0.0028 0.0014c0.3564 0.1469 0.6537 0.3543 0.9047 0.63l0.0018 0.0021 0.0017 0.0021c0.2596 0.2765 0.462 0.611 0.608 1.0157l0.001 0.0014v0.0011c0.145 0.3921 0.219 0.8331 0.219 1.3295 0 0.5057-0.074 0.9514-0.219 1.3428-0.146 0.396-0.3483 0.7272-0.609 1.0049l-0.0017 0.0021-0.0018 0.0021c-0.2514 0.2762-0.5505 0.4887-0.9082 0.6443-0.3557 0.1461-0.7435 0.2206-1.1743 0.2206-0.4309 0-0.8189-0.0744-1.1747-0.2206-0.3576-0.1556-0.6618-0.3693-0.9228-0.6467-0.2521-0.2776-0.4514-0.6097-0.5978-1.0067-0.1443-0.3914-0.219-0.8371-0.219-1.3428 0-0.4964 0.0744-0.9374 0.219-1.3295l0.0003-0.0011 0.0007-0.0014c0.1467-0.4057 0.3458-0.7411 0.5968-1.0174 0.2605-0.277 0.5633-0.4856 0.9197-0.6325l0.0028-0.0014 0.0031-0.0014c0.3553-0.1542 0.742-0.2325 1.1719-0.2325z"></path></g></g></svg>'
  }
});
