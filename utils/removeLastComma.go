package utils


/**
 * Remove the last comma from the string
 * (usually used for patch insert queries)
 */
func RemoveLastComma(query string) string  {
  queryx := len(query)

  if queryx > 0 && query[queryx-1] == ',' {
    query = query[:queryx-1]
  }

  return query
}
