package review

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Sanitize(t *testing.T) {
	r := require.New(t)

	tcases := []struct {
		input string
		output string
	}{
		{
			`<option value="<%= state.Code %>" <%= if (selectedState == state.Code) { %> selected <% } %>> <%= state.Name %></option>`,
			`<option value="ERB" ERB selected ERB> ERB</option>`,
		},
		{
			`    
			<div class="listing-header">
				<div class="row">
					<div class="col-md-6">
						<%= linkTo(backURL()) { %>
							<h4 class=""><i class="fa fa-arrow-left"></i>&nbsp;<strong>Edit Invoice</strong></h4>
						<% } %>
					</div>
				</div>
			</div>`,
			`    
			<div class="listing-header">
				<div class="row">
					<div class="col-md-6">
						ERB
							<h4 class=""><i class="fa fa-arrow-left"></i>&nbsp;<strong>Edit Invoice</strong></h4>
						ERB
					</div>
				</div>
			</div>`,
		},
		{
			` 
			<%= form_for(invoice, {action: "/invoices/"+invoice.ID, method: "PUT", class: "new-form col-md-12"}) { %>
                <%= partial("invoices/form.plush.html") %>

                <div class="col-md-12">
                    <button class="success-button no-radius pull-right">Update Invoice</button>
                </div>
			<% } %>`,
			` 
			ERB
                ERB

                <div class="col-md-12">
                    <button class="success-button no-radius pull-right">Update Invoice</button>
                </div>
			ERB`,
		},
	}

	for _, tcase := range tcases {
		output := sanitizeERB([]byte(tcase.input))
		
		r.Equal(tcase.output, string(output))
	}
}