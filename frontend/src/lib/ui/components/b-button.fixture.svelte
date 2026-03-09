<script lang="ts">
  import { BButton } from "$lib/ui";
</script>

<div class="grid gap-2">
  <BButton variant="primary" leadingIcon="addLogEvent">Primary Action</BButton>
  <BButton variant="secondary" trailingIcon="editAnimal" disabled={true}>
    Disabled Secondary
  </BButton>
</div>
