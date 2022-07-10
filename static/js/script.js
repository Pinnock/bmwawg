let attention = Prompt();

function notifyModal(title, text, icon, confirmBtnTxt) {
    Swal.fire({
        title: title,
        text: text,
        icon: icon,
        confirmButtonText: confirmBtnTxt
    })
}

function Prompt() {
    let toast = function (c) {
        const {
            msg = "",
            icon = "success",
            position = "center-top",
        } = c;

        const Toast = Swal.mixin({
            toast: true,
            title: msg,
            position: position,
            icon: icon,
            showConfirmButton: false,
            timer: 3000,
            timerProgressBar: true,
            didOpen: (toast) => {
                toast.addEventListener('mouseenter', Swal.stopTimer)
                toast.addEventListener('mouseleave', Swal.resumeTimer)
            }
        })

        Toast.fire({})
    }

    let success = function (c) {
        const {
            msg = "",
            title = "",
            footer = "",
        } = c;

        Swal.fire({
            icon: 'success',
            title: title,
            text: msg,
            footer: footer,
        })
    }

    let error = function (c) {
        const {
            msg = "",
            title = "",
            footer = "",
        } = c;

        Swal.fire({
            icon: 'error',
            title: title,
            text: msg,
            footer: footer,
        })
    }

    async function custom(c) {
        const {
            msg = "",
            title = "",
            backdrop = true
        } = c;

        const { value: result } = await Swal.fire({
            title: title,
            html: msg,
            backdrop: backdrop,
            focusConfirm: false,
            showCancelButton: true,
            willOpen: () => {
                const resDates = document.getElementById("reservation-dates-modal")
                const datePicker = new DateRangePicker(
                    resDates,
                    { format: "mm/dd/yyyy", showOnFocus: true, orientation: "top" },
                )
            },
            didOpen: () => {
                document.getElementById("start").removeAttribute("disabled");
                document.getElementById("end").removeAttribute("disabled");
            },
            preConfirm: () => {
                return [
                    document.getElementById('start').value,
                    document.getElementById('end').value
                ]
            }
        })

        if (result && result.dismiss !== Swal.DismissReason.cancel && 
            result.value !== "" && c.callback !== undefined){
                c.callback(result);
        } 
    }
    return {
        toast: toast,
        success: success,
        error: error,
        custom: custom
    }
};


(function () {
    'use strict'

    // Fetch all the forms we want to apply custom Bootstrap validation styles to
    var forms = document.querySelectorAll('.needs-validation')

    // Loop over them and prevent submission
    Array.prototype.slice.call(forms)
        .forEach(function (form) {
            form.addEventListener('submit', function (event) {
                if (!form.checkValidity()) {
                    event.preventDefault()
                    event.stopPropagation()
                }

                form.classList.add('was-validated')
            }, false)
        })
})()
