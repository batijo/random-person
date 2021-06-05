<template>
    <header>
        <h1>{{title}}</h1>
    </header>
</template>

<script>
export default {
    name: "Header",
    props: {
        title: String,
    }
}
</script>