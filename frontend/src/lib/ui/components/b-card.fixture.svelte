<script lang="ts">
  import { BCard } from "$lib/ui";
</script>

<BCard class="fixture-structured">
  {#snippet header()}
    <h3>Card Header</h3>
  {/snippet}

  {#snippet body()}
    <p>Body Snippet</p>
  {/snippet}

  {#snippet footer()}
    <button type="button">Footer Action</button>
  {/snippet}

  <p>Should not render</p>
</BCard>

<BCard class="fixture-children mt-2">
  <p>Children Content</p>
</BCard>
