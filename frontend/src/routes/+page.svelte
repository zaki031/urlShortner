<script>
  const apiUrl = import.meta.env.VITE_API_URL;
  let url = "";
  let result = "";
  let errorMessage = "";

  async function submitForm(event) {
    event.preventDefault();


    try {
      const res = await fetch(apiUrl, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ url: url }),
      });
      if (res.status == 400) {
        throw new Error("Host not found");
      }
      let resText = await res.text();
      const json = JSON.parse(resText);
      result = JSON.stringify(json, null);
    } catch (error) {
      errorMessage = error;
    }
  }
</script>

<svelte:head>
  <title>URL Shortener</title>
  <meta name="description" content="URL Shortener App" />
</svelte:head>

<section>
  <form on:submit={submitForm}>
    <input
      name="url"
      placeholder="Enter the URL you want to shorten"
      bind:value={url}
      type="text"
      required
    />
    <button type="submit">Submit</button>
  </form>
  {#if result}
    <p>{result}</p>
  {/if}
  {#if errorMessage}
    <p style="color: red;">{errorMessage}</p>
  {/if}
</section>

<style>
  input,
  form {
    width: 100%;
  }
  section {
    padding: 20px;
  }
</style>
