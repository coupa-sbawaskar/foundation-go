package rest

func ConvertHttpQueryParamsToPersistenceParams(c *gin.Context) persistence.QueryParams {
	//converts http query params into persistence.QueryParams
	//the http query params can be quite complex depending on if nested logical expressions are used. e.g,
	//path?name=foo  => "where name='foo'"
	//path?name=foo|bar => "where name='foo' or name='bar'"

	//there are so many ways to do the below http query, below is just an example thoughtup on the fly. all apis do this differently
	//path?:OR(name=foo & bar) => "where (first_name='first1' and last_name='last1') or  (first_name='first2' and last_name='last2')"
	//whatever format/syntax used should be documented in https://coupadev.atlassian.net/wiki/spaces/CPL/pages/543949398/Microservice+Standards+Document

	return persistence.QueryParams{...}
}
