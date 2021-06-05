<template>
    <div>
        <h2>{{person.name}} {{person.surname}}</h2>
    </div>
</template>

<script>
export default {
    name: "Person",
    props: {
        person: Object,
    },
}
</script>