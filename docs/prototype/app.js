(function () {
  function initIcons() {
    if (window.lucide && typeof window.lucide.createIcons === 'function') {
      window.lucide.createIcons();
    }
  }

  function setChipState(chip, isActive) {
    chip.classList.toggle('bg-emerald-700', isActive);
    chip.classList.toggle('text-white', isActive);
    chip.classList.toggle('border-emerald-700', isActive);

    chip.classList.toggle('bg-white', !isActive);
    chip.classList.toggle('text-stone-700', !isActive);
    chip.classList.toggle('border-stone-300', !isActive);
    chip.setAttribute('aria-pressed', isActive ? 'true' : 'false');
  }

  function initHomeTagFilter() {
    var chips = document.querySelectorAll('.chip-filter[data-filter]');
    var cards = document.querySelectorAll('.animal-card[data-tag]');
    if (!chips.length || !cards.length) return;

    function setFilter(filter) {
      chips.forEach(function (chip) {
        var isActive = chip.getAttribute('data-filter') === filter;
        setChipState(chip, isActive);
      });

      cards.forEach(function (card) {
        var tag = card.getAttribute('data-tag');
        var show = filter === 'all' || tag === filter;
        card.classList.toggle('hidden', !show);
        card.classList.toggle('block', show);
      });
    }

    chips.forEach(function (chip) {
      chip.addEventListener('click', function () {
        setFilter(chip.getAttribute('data-filter'));
      });
    });

    setFilter('all');
  }

  function setEventTypeStyle(activeType) {
    var options = document.querySelectorAll('.event-type-option[data-event-type]');
    options.forEach(function (option) {
      var isActive = option.getAttribute('data-event-type') === activeType;

      option.classList.toggle('bg-emerald-700', isActive);
      option.classList.toggle('text-white', isActive);
      option.classList.toggle('border-emerald-700', isActive);

      option.classList.toggle('bg-white', !isActive);
      option.classList.toggle('text-stone-700', !isActive);
      option.classList.toggle('border-stone-300', !isActive);
    });
  }

  function initEventTypeForm() {
    var eventTypeRadios = document.querySelectorAll('input[name="event-type"]');
    if (!eventTypeRadios.length) return;

    var groups = {
      feed: ['amount'],
      medication: ['med-type', 'dosage'],
      weight: ['weight-value'],
      note: ['note-text']
    };

    function setVisible(type) {
      var allIds = ['amount', 'med-type', 'dosage', 'weight-value', 'note-text'];
      allIds.forEach(function (id) {
        var input = document.getElementById(id);
        if (!input) return;
        var field = input.closest('.field');
        if (!field) return;
        field.style.display = groups[type].indexOf(id) > -1 ? 'block' : 'none';
      });
      setEventTypeStyle(type);
    }

    eventTypeRadios.forEach(function (radio) {
      radio.addEventListener('change', function () {
        if (radio.checked) setVisible(radio.id);
      });
    });

    var checked = document.querySelector('input[name="event-type"]:checked');
    setVisible(checked ? checked.id : 'feed');
  }

  initIcons();
  initHomeTagFilter();
  initEventTypeForm();
})();
