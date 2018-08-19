$(document).ready(function(){
	$('#ajax_id').click(function (){
		$("#result").html("loading")
		var dataForm = $('#form-login').serialize();
		$.ajax({
			url : "/login",
			type : "POST",
			data : dataForm,
			success: function(result){
				$("#result").html("sukses")
			}
		})
	})
})