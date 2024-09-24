<script>
    const apiUrl = import.meta.env.VITE_API_URL;

    import { page } from '$app/stores';
    import { onMount } from 'svelte';

    function redirect(url) {
        window.location.replace(url);
    }

    let data = [];
    let error = null;
    let isLoading = true;
    let slug = '';

    $: slug = $page.params.slug;

    onMount(async () => {
        try {
            const res = await fetch(apiUrl + slug, {
                headers: {
                    'Content-Type': 'application/json',
                }
            });

            if (!res.ok) {
                error = res.status;
            }
            data = await res.json();
            redirect(data);
       
        } catch (err) {
            console.error('Error:', err); 
        } finally {
            isLoading = false;
        }
    });
</script>

<svelte:head>
    <title>Home</title>
    <meta name="description" content="" />
</svelte:head>

<section>
    
    {#if isLoading}
        <p>Loading...</p>
    {:else if error=== 404 }
        <p>Short URL doesn't exist.</p>
    {:else}
        <p>You are being redirected to the original website.</p>
    {/if}
</section>
