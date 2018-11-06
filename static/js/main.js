
$(document).ready(function () {
    $.ajax({
        url: '/hello',
        type: "get"
    }).done(
        function (data) {
            $('#content').append(data);
        });
});
