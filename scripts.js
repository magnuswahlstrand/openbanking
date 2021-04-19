
var cardTemplate = `<div class="card">
<div class="card-header p-0" id="heading_{{index}}">
  <div
    class="row p-2"
    data-toggle="collapse"
    data-target="#collapse_{{index}}"
    aria-expanded="true"
    aria-controls="collapse_{{index}}"
  >
    <div class="col-1">
      <svg
        height="48"
        class="octicon octicon-repo pull-left"
        viewBox="0 0 12 16"
        version="1.1"
        width="36"
        aria-hidden="true"
        fill="purple"
      >
        <path
          fill="lightgray"
          fill-rule="evenodd"
          d="M4 9H3V8h1v1zm0-3H3v1h1V6zm0-2H3v1h1V4zm0-2H3v1h1V2zm8-1v12c0 .55-.45 1-1 1H6v2l-1.5-1.5L3 16v-2H1c-.55 0-1-.45-1-1V1c0-.55.45-1 1-1h10c.55 0 1 .45 1 1zm-1 10H1v2h2v-1h3v1h5v-2zm0-10H2v9h9V1z"
        ></path>
      </svg>
    </div>

    <div class="col-5  d-flex align-items-center">
        <div class="d-flex align-items-start flex-column bd-highlight">
            <div class="mb-auto"><b>{{type}}</b></div>
            <div style="color:gray">{{bban}}</span></div>
        </div>
    </div>

    <div class="col-5  d-flex  align-items-center ml-auto" style="font-size:1.5em; color: gray;">
        <div class="ml-auto p-2 bd-highlight">{{print_balance}}</div>
    </div>

    <div class="col-1">
    <div class="dropright">
    <button
      class="btn test"
      href="#"
      role="button"
      id="dropdownMenu1"
      data-toggle="dropdown"
      aria-haspopup="true"
      aria-expanded="true"
    >
    </button>

    <div class="dropdown-menu" aria-labelledby="dropdownMenu1" aria-expanded="true">
      <a class="dropdown-item" href="#" onclick="copyToClipboard(event, '{{ available_balance }}', 'balance')">Copy balance</a>
      <a class="dropdown-item" href="#" onclick="copyToClipboard(event, '{{ bban }}', 'BBAN')">Copy BBAN</a>
      <a class="dropdown-item" href="#" onclick="copyToClipboard(event, '{{ iban }}', 'IBAN')">Copy IBAN</a>
    </div>
    </div>

    </div>    
  </div>
</div>

<div
  id="collapse_{{index}}"
  class="collapse"
  aria-labelledby="heading_{{index}}"
  data-parent="#accordionAccounts"
>
  <div class="card-body">
    {{{metadata}}}
  </div>
</div>
</div>`;

function generateCard(account, index) {
  account["index"] = index;
  account["print_balance"] = account["available_balance"].toFixed(2);

  var rendered = Mustache.render(cardTemplate, account);
  return rendered;
}

function copyToClipboard(event, value, alertName) {
  updateClipboard(value, alertName);
  event.preventDefault();
  event.stopPropagation();
}

// Show a simple toast message that fades automatically
// Element is removed afterwards
function createToast(message) {
  var elem = $('' +
      '<div class="toast" aria-live="assertive" aria-atomic="true">' +
      '<div class="toast-body">' +
      message +
      '</div>' +
      '</div>')

  $(elem).toast('show');

  // Remove element after it has been hidden
  $(elem).on('hidden.bs.toast', function () {
    $(elem).remove();
  })

  $("body").append(elem);
}

function updateClipboard(value, alertName) {
  navigator.clipboard.writeText(value).then(
    function() {
      console.log("copied", value);
      alertName = alertName.charAt(0).toUpperCase() +
      alertName.slice(1);

      createToast(alertName + " copied.")
    },
    function() {}
  );
}
